package server

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (*minionServer_s) Heartbeat(ctx context.Context, req *proto.HeartbeatRequest) (*proto.HeartbeatResponse, error) {
	logrus.WithField("minion_id", req.GetMinionId()).Info("received heartbeat")

	return &proto.HeartbeatResponse{}, nil
}
