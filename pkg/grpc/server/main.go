package server

import (
	"context"
	"fmt"
	"net"

	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type minionServer_s struct {
	proto.UnimplementedMinionServiceServer

	ScoreTasks         <-chan *proto.GetScoreTaskResponse
	ScoreTaskResponses chan<- *proto.SubmitScoreTaskRequest
}

var (
	grpcServer   *grpc.Server
	minionServer *minionServer_s
)

func Serve(ctx context.Context, scoreTaskChan <-chan *proto.GetScoreTaskResponse, scoreTaskReponseChan chan<- *proto.SubmitScoreTaskRequest) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GRPC.Host, config.GRPC.Port))
	if err != nil {
		logrus.WithError(err).Fatal("encountered error while starting gRPC server")
	}

	// TODO: Implement TLS Configuration

	grpcServer = grpc.NewServer()

	minionServer = &minionServer_s{
		ScoreTasks:         scoreTaskChan,
		ScoreTaskResponses: scoreTaskReponseChan,
	}

	proto.RegisterMinionServiceServer(grpcServer, minionServer)

	logrus.Infof("gRPC server listening on %s:%d", config.GRPC.Host, config.GRPC.Port)

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
	}()

	err = grpcServer.Serve(lis)
	if err != nil && err != grpc.ErrServerStopped {
		logrus.WithError(err).Fatal("encountered error while serving gRPC server")
	} else if err == grpc.ErrServerStopped {
		logrus.Info("gRPC server stopped")
	}
}
