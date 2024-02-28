package engine

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/scorify/backend/pkg/structs"
	"github.com/sirupsen/logrus"
)

type EngineState int

const (
	Stopped EngineState = iota
	Running
)

type engine struct {
	lock        *structs.Lock
	ctx         context.Context
	ent         *ent.Client
	redis       *redis.Client
	taskChan    chan<- *proto.GetScoreTaskResponse
	resultsChan <-chan *proto.SubmitScoreTaskRequest
}

func NewEngine(
	ctx context.Context,
	entClient *ent.Client,
	redis *redis.Client,
	taskChan chan<- *proto.GetScoreTaskResponse,
	resultsChan <-chan *proto.SubmitScoreTaskRequest,
) *engine {
	return &engine{
		lock:        structs.NewLock(),
		ctx:         ctx,
		ent:         entClient,
		redis:       redis,
		taskChan:    taskChan,
		resultsChan: resultsChan,
	}
}

func (e *engine) Stop() error {
	err := e.lock.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (e *engine) Start() error {
	err := e.lock.Lock()
	if err != nil {
		return err
	}

	go e.loop()

	return nil
}

func (e *engine) IsStopped() bool {
	return e.lock.IsUnlocked()
}

func (e *engine) IsRunning() bool {
	return e.lock.IsLocked()
}

func (e *engine) State() EngineState {
	if e.IsStopped() {
		return Stopped
	}

	return Running
}

func (e *engine) loop() {
	ticker := time.NewTicker(config.Interval)

	defer ticker.Stop()
	defer e.lock.Unlock()

	for {
		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			err := e.loopRoundRunner()
			if err != nil {
				logrus.WithError(err).Error("failed to run round")
				return
			}
		}
	}
}

func (e *engine) loopRoundRunner() error {
	err := e.lock.Lock()
	if err != nil {
		return err
	}

	roundCtx, cancel := context.WithTimeout(e.ctx, config.Interval)
	defer cancel()

	// Get the current round number
	var roundNumber int
	entLastRound, err := e.ent.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		Limit(1).
		Select(round.FieldNumber).
		Only(e.ctx)
	if err != nil {
		roundNumber = 1
	} else {
		roundNumber = entLastRound.Number + 1
	}

	// Create new round
	entRound, err := e.ent.Round.Create().
		SetNumber(roundNumber).
		Save(e.ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to create new round")
		return nil
	}

	err = e.lock.Lock()
	if err != nil {
		return err
	}

	// Run round
	return e.runRound(roundCtx, entRound)
}

func (e *engine) runRound(ctx context.Context, entRound *ent.Round) error {
	defer e.lock.Unlock()

	// Get all the tasks
	tasks, err := e.ent.CheckConfig.Query().
		WithUser().
		WithCheck().
		Order(
			// ID are uuids and thus check orders are shuffled
			ent.Desc(checkconfig.FieldID),
		).
		All(ctx)
	if err != nil {
		return err
	}

	// Bulk create tasks
	entStatusCreates := make([]*ent.StatusCreate, len(tasks))
	for i, task := range tasks {
		entStatusCreates[i] = e.ent.Status.Create().
			SetRound(entRound).
			SetUser(task.Edges.User).
			SetCheck(task.Edges.Check).
			SetPoints(task.Edges.Check.Weight).
			SetStatus(status.StatusUnknown)
	}

	entStatuses, err := e.ent.Status.CreateBulk(entStatusCreates...).Save(ctx)
	if err != nil {
		return err
	}

	// Create a map of round tasks to keep track of the tasks
	roundTasks := make(map[uuid.UUID]bool)

	for _, entStatus := range entStatuses {
		roundTasks[entStatus.ID] = true
	}

	// Submit tasks to the workers
	go func() {
		for _, entStatus := range entStatuses {
			conf, err := json.Marshal(entStatus.Edges.Check.Config)
			if err != nil {
				logrus.WithError(err).Error("failed to marshal check config")
				continue
			}

			e.taskChan <- &proto.GetScoreTaskResponse{
				StatusId:   entStatus.ID.String(),
				SourceName: entStatus.Edges.Check.Source,
				Config:     string(conf),
			}
		}
	}()

	// Wait for the results
	for len(roundTasks) > 0 {
		select {
		case <-ctx.Done():
			return nil
		case result := <-e.resultsChan:
			status_id, err := uuid.Parse(result.StatusId)
			if err != nil {
				logrus.WithError(err).Error("failed to parse status id")
				continue
			}

			go e.updateStatus(ctx, roundTasks, status_id, result.Error, result.Status)
		}
	}

	for status_id := range roundTasks {
		go func(id uuid.UUID) {
			_, err := e.ent.Status.UpdateOneID(id).
				SetStatus(status.StatusUnknown).
				SetPoints(0).
				Save(ctx)
			if err != nil {
				logrus.WithField("id", id).WithError(err).Error("failed to update status")
			}
		}(status_id)
	}

	return nil
}

func (e *engine) updateStatus(ctx context.Context, roundTasks map[uuid.UUID]bool, status_id uuid.UUID, errorMessage string, _status proto.Status) {
	_, ok := roundTasks[status_id]
	if !ok {
		logrus.WithField("status_id", status_id).Error("uuid not belong to round was submitted")
		return
	}

	entStatusUpdate := e.ent.Status.UpdateOneID(status_id).
		SetStatus(status.Status(_status))

	if errorMessage != "" {
		entStatusUpdate.SetError(errorMessage)
	}

	if _status != proto.Status_up {
		entStatusUpdate.SetPoints(0)
	}

	_, err := entStatusUpdate.Save(ctx)
	if err != nil {
		logrus.WithField("id", status_id).WithError(err).Error("failed to update status")
		return
	}

	delete(roundTasks, status_id)
}
