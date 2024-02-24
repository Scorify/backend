package grpc

import (
	"github.com/scorify/backend/pkg/config"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "grpc",
	Short:   "gRPC subcommands",
	Long:    "gRPC subcommands",
	Aliases: []string{"g"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(
		clientCmd,
		serverCmd,
	)
}
