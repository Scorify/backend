package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/scorify/backend/pkg/structs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type minionServer_s struct {
	proto.UnimplementedMinionServiceServer

	counter            *structs.Counter
	ScoreTasks         <-chan *proto.GetScoreTaskResponse
	ScoreTaskResponses chan<- *proto.SubmitScoreTaskRequest
}

var (
	grpcServer   *grpc.Server
	minionServer *minionServer_s
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	var minionID string
	switch req := req.(type) {
	case *proto.GetScoreTaskRequest:
		minionServer.counter.Increment()
		defer minionServer.counter.Decrement()
		minionID = req.MinionId
	case *proto.SubmitScoreTaskRequest:
		minionID = req.MinionId
	case *proto.HeartbeatRequest:
		minionID = req.MinionId
	default:
		minionID = "unknown"
	}

	resp, err := handler(ctx, req)

	logrus.WithFields(logrus.Fields{
		"method":   info.FullMethod,
		"took":     time.Since(start),
		"minions":  minionServer.counter.Get(),
		"minionID": minionID,
	}).Debug("gRPC request")

	return resp, err
}

func Serve(ctx context.Context, scoreTaskChan <-chan *proto.GetScoreTaskResponse, scoreTaskReponseChan chan<- *proto.SubmitScoreTaskRequest) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GRPC.Host, config.GRPC.Port))
	if err != nil {
		logrus.WithError(err).Fatal("encountered error while starting gRPC server")
	}

	// TODO: Implement TLS Configuration

	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	minionServer = &minionServer_s{
		ScoreTasks:         scoreTaskChan,
		ScoreTaskResponses: scoreTaskReponseChan,
		counter:            structs.NewCounter(),
	}

	proto.RegisterMinionServiceServer(grpcServer, minionServer)

	logrus.Infof("Starting gRPC server on %s:%d", config.GRPC.Host, config.GRPC.Port)

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
