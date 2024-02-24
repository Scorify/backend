package client

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn   *grpc.ClientConn
	Client proto.ScorifyClient
)

func Open() {
	ctx := context.Background()

	_conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to gRPC server")
	}

	conn = _conn

	Client = proto.NewScorifyClient(conn)
}

func Close() {
	err := conn.Close()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to close gRPC connection")
	}
}
