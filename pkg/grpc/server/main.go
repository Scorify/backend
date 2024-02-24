package server

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type scorifyServer_s struct {
	proto.UnimplementedScorifyServer
}

var (
	grpcServer    *grpc.Server
	scorifyServer scorifyServer_s = scorifyServer_s{}

	send_chan = make(chan *proto.CheckRequest)
	recv_chan = make(chan *proto.CheckResponse)
)

func Serve(ctx context.Context) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		logrus.WithError(err).Fatal("gRPC server failed to listen")
	}

	grpcServer = grpc.NewServer()

	proto.RegisterScorifyServer(grpcServer, &scorifyServer)

	logrus.Infof("gRPC server is listening on port %d", config.GRPC.Port)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
		wg.Done()
	}()

	go func() {
		readerWriter := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

		go func() {
			for resp := range recv_chan {
				fmt.Println("X-Recieved: ", resp)
			}
		}()

		for {
			fmt.Print("> ")
			text, err := readerWriter.ReadString('\n')
			if err != nil {
				logrus.WithError(err).Error("failed to read from stdin")
				continue
			}

			switch text {
			case "exit\n":
				ctx.Done()
				return
			default:
				send_chan <- &proto.CheckRequest{
					Check:  text,
					Config: text,
				}
				fmt.Println("X-Sent: ", text)
			}
		}
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.WithError(err).Fatal("gRPC server failed to serve")
	}

	wg.Wait()

	logrus.Info("gRPC server is closed")
}
