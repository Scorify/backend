package grpc

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the gRPC server",
	Long:    "Run the gRPC server",
	Aliases: []string{"s"},

	Run: serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {
	server.Serve(context.Background())
}
