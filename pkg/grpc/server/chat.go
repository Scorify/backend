package server

import (
	"fmt"

	"github.com/scorify/backend/pkg/grpc/proto"
)

func (scorifyServer_s) Check(stream proto.Scorify_CheckServer) error {
	err_chan := make(chan error)

	go func() {
		for {
			req, err := stream.Recv()
			fmt.Println("Recieved: ", req)
			if err != nil {
				err_chan <- err
				return
			}

			recv_chan <- req
		}
	}()

	go func() {
		for send := range send_chan {
			fmt.Println("Sending: ", send)
			err := stream.Send(send)
			if err != nil {
				err_chan <- err
				return
			}

			send_chan <- send
		}
	}()

	err := <-err_chan
	fmt.Println("Error: ", err)
	return err
}
