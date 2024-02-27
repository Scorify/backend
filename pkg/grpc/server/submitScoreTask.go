package server

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/proto"
)

func (s *minionServer_s) SubmitScoreTask(ctx context.Context, req *proto.SubmitScoreTaskRequest) (*proto.SubmitScoreTaskResponse, error) {
	s.ScoreTaskResponses <- req

	return &proto.SubmitScoreTaskResponse{}, nil
}
