package server

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (s *minionServer_s) SubmitScoreTask(ctx context.Context, req *proto.SubmitScoreTaskRequest) (*proto.SubmitScoreTaskResponse, error) {
	logrus.WithField("minion_id", req.GetMinionId()).Info("received submit score task request")

	s.ScoreTaskResponses <- req

	return &proto.SubmitScoreTaskResponse{}, nil
}
