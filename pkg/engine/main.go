package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/cache"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/graph/model"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/scorify/backend/pkg/structs"
	"github.com/sirupsen/logrus"
)

type state int

const (
	EnginePaused state = iota
	EngineWaiting
	EngineRunning
)

type Client struct {
	lock        *sync.Mutex
	state       state
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
) *Client {
	return &Client{
		lock:        &sync.Mutex{},
		state:       EnginePaused,
		ctx:         ctx,
		ent:         entClient,
		redis:       redis,
		taskChan:    taskChan,
		resultsChan: resultsChan,
	}
}

func (e *Client) Stop() error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.state == EngineWaiting {
		e.state = EnginePaused
		_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStatePaused)
		return err
	}

	return fmt.Errorf("cannot stop engine from state %q", e.state)
}

func (e *Client) Start() error {
	if e.state == EngineRunning {
		return fmt.Errorf("engine already running")
	}

	go e.loop()

	e.state = EngineWaiting
	_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStateWaiting)
	return err
}

func (e *Client) State() (model.EngineState, error) {
	switch e.state {
	case EnginePaused:
		return model.EngineStatePaused, nil
	case EngineWaiting:
		return model.EngineStateWaiting, nil
	case EngineRunning:
		return model.EngineStateRunning, nil
	}

	return "", fmt.Errorf("unknown engine state %q", e.state)
}

func (e *Client) loop() {
	ticker := time.NewTicker(config.Interval)

	defer func() {
		ticker.Stop()
		e.state = EnginePaused
		cache.PublishEngineState(context.Background(), e.redis, model.EngineStatePaused)
	}()

	for {
		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			if e.state == EnginePaused {
				return
			}

			err := e.loopRoundRunner()
			if err != nil {
				logrus.WithError(err).Error("failed to run round")
				return
			}

			e.state = EngineWaiting
			_, err = cache.PublishEngineState(context.Background(), e.redis, model.EngineStateWaiting)
			if err != nil {
				logrus.WithError(err).Error("failed to publish engine state")
				return
			}
		}
	}
}

func (e *Client) loopRoundRunner() error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.state = EngineRunning
	_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStateRunning)
	if err != nil {
		return err
	}

	roundCtx, cancel := context.WithTimeout(e.ctx, config.Interval-time.Millisecond*500)
	defer cancel()

	// Get the current round number
	var roundNumber int
	entLastRound, err := e.ent.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		First(e.ctx)
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

	logrus.WithField("time", time.Now()).Infof("Running round %d", roundNumber)

	// Run round
	return e.runRound(roundCtx, entRound)
}

func (e *Client) runRound(ctx context.Context, entRound *ent.Round) error {
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
	roundTasks := structs.NewSyncMap[uuid.UUID, *ent.CheckConfig]()

	for i, entStatus := range entStatuses {
		roundTasks.Set(entStatus.ID, tasks[i])
	}

	// Submit tasks to the workers
	go func() {
		for _, entStatus := range entStatuses {
			entConfig, ok := roundTasks.Get(entStatus.ID)
			if !ok {
				logrus.WithField("id", entStatus.ID).Error("failed to get task")
				continue
			}

			conf, err := json.Marshal(entConfig.Config)
			if err != nil {
				logrus.WithError(err).Error("failed to marshal check config")
				continue
			}

			e.taskChan <- &proto.GetScoreTaskResponse{
				StatusId:   entStatus.ID.String(),
				SourceName: entConfig.Edges.Check.Source,
				Config:     string(conf),
			}
		}
	}()

	// Wait for the results
	for roundTasks.Legnth() > 0 {
		select {
		case <-ctx.Done():
			return nil
		case result := <-e.resultsChan:
			status_id, err := uuid.Parse(result.StatusId)
			if err != nil {
				logrus.WithError(err).Error("failed to parse status id")
				continue
			}

			switch result.Status {
			case proto.Status_up:
				go e.updateStatus(ctx, roundTasks, status_id, result.Error, status.StatusUp)
			case proto.Status_down:
				go e.updateStatus(ctx, roundTasks, status_id, result.Error, status.StatusDown)
			case proto.Status_unknown:
				go e.updateStatus(ctx, roundTasks, status_id, result.Error, status.StatusUnknown)
			default:
				go e.updateStatus(ctx, roundTasks, status_id, result.Error, status.StatusUnknown)
				logrus.WithFields(logrus.Fields{
					"status":    result.Status,
					"status_id": status_id,
				}).Error("unknown status")
			}
		}
	}

	for status_id := range roundTasks.Map() {
		entStatus, err := e.ent.Status.UpdateOneID(status_id).
			SetStatus(status.StatusUnknown).
			SetPoints(0).
			Save(ctx)
		if err != nil {
			logrus.WithField("id", status_id).WithError(err).Error("failed to update status")
		} else {
			logrus.WithField("status", entStatus).Info("status not reported, set to 0")
		}

		_, err = cache.PublishScoreStream(ctx, e.redis, entStatus)
		if err != nil {
			logrus.WithError(err).Error("failed to publish score stream")
		}
	}

	var users []struct {
		UserID uuid.UUID `json:"user_id"`
		Points int       `json:"points"`
	}

	err = e.ent.Status.Query().
		Where(
			status.HasRoundWith(round.ID(entRound.ID)),
		).
		GroupBy(status.FieldUserID).
		Aggregate(ent.Sum(status.FieldPoints)).
		Scan(ctx, &users)
	if err != nil {
		logrus.WithError(err).Error("failed to aggregate points")
		return err
	}

	entScoreCacheCreates := make([]*ent.ScoreCacheCreate, len(users))
	for i, user := range users {
		entScoreCacheCreates[i] = e.ent.ScoreCache.Create().
			SetRound(entRound).
			SetUserID(user.UserID).
			SetPoints(user.Points)
	}

	_, err = e.ent.ScoreCache.CreateBulk(entScoreCacheCreates...).Save(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to create score cache")
	}
	logrus.Warn("Round finished, waiting for next round")
	return err
}

func (e *Client) updateStatus(ctx context.Context, roundTasks *structs.SyncMap[uuid.UUID, *ent.CheckConfig], status_id uuid.UUID, errorMessage string, _status status.Status) {
	_, ok := roundTasks.Get(status_id)
	if !ok {
		logrus.WithField("status_id", status_id).Error("uuid not belong to round was submitted")
		return
	}

	entStatusUpdate := e.ent.Status.UpdateOneID(status_id).
		SetStatus(status.Status(_status))

	if errorMessage != "" {
		entStatusUpdate.SetError(errorMessage)
	}

	if _status != status.StatusUp {
		entStatusUpdate.SetPoints(0)
	}

	entStatus, err := entStatusUpdate.Save(ctx)
	if err != nil {
		logrus.WithField("id", status_id).WithError(err).Error("failed to update status")
		return
	}

	logrus.WithField("status", entStatus).Info("status updated")

	roundTasks.Delete(status_id)

	_, err = cache.PublishScoreStream(ctx, e.redis, entStatus)
	if err != nil {
		logrus.WithError(err).Error("failed to publish score stream")
	}
}
