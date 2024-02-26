package server

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (s *minionServer_s) GetScoreTask(ctx context.Context, req *proto.GetScoreTaskRequest) (*proto.GetScoreTaskResponse, error) {
	logrus.WithField("minion_id", req.GetMinionId()).Info("received get score task request")

	return <-s.ScoreTasks, nil
}
