package minion

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "minion",
	Short:   "Start scoring minion worker",
	Long:    "Start scoring minion worker",
	Aliases: []string{"m", "worker", "w"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	grpcClient, err := client.Open(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("encountered error while opening gRPC client")
	}

	defer grpcClient.Close()

	logrus.Info("gRPC client opened successfully")

	go func() {
		for {
			grpcClient.Heartbeat(ctx)
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		task, err := grpcClient.GetScoreTask(ctx)
		if err != nil {
			logrus.WithError(err).Fatal("encountered error while getting score task")
		}

		logrus.WithField("task", task).
			WithField("time", time.Now()).
			Info("recieved score task")

		uuid, err := uuid.Parse(task.GetStatusId())
		if err != nil {
			logrus.WithError(err).Fatal("encountered error while parsing UUID")
		}

		_, err = grpcClient.SubmitScoreTask(ctx, uuid, "bruh", status.StatusUp)
		if err != nil {
			logrus.WithError(err).Fatal("encountered error while submitting score task")
		}
	}
}
