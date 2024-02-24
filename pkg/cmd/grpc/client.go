package grpc

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/scorify/backend/pkg/grpc/client"
	"github.com/scorify/backend/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:     "client",
	Short:   "Run the gRPC client",
	Long:    "Run the gRPC client",
	Aliases: []string{"c"},

	Run: clientRun,
}

func clientRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	stream, err := client.Client.Check(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to check gRPC server")
	}

	err_chan := make(chan error)

	go func() {
		defer fmt.Println("Closing")
		for {
			fmt.Println("waiting for message")
			req, err := stream.Recv()
			fmt.Println("Recieved: ", req)
			if err != nil {
				fmt.Println("Error: ", err)
				err_chan <- err
				return
			}

			fmt.Println("Recieved: ", req, "\n> ")
		}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				err_chan <- err
				return
			}
			switch text {
			case "exit":
				err_chan <- nil
				return
			default:
				err = stream.Send(&proto.CheckResponse{Error: text})
				if err != nil {
					err_chan <- err
					return
				}
			}
		}
	}()

	err = <-err_chan
	if err != nil {
		logrus.WithError(err).Fatal("Failed to communicate with gRPC server")
	}
}
