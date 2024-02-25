package minion

import (
	"context"
	"fmt"

	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/graph/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "minion",
	Short:   "Run the minion",
	Long:    "Run the minion",
	Aliases: []string{"m"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	gqlClient, err := client.NewClient()
	if err != nil {
		logrus.WithError(err).Fatal("error while creating gql client")
	}

	fmt.Println("Enrolling minion")

	fmt.Println(gqlClient.EnrollMinion(context.Background()))
	// err = gqlClient.SubmitStatus(context.Background(), uuid.New(), status.StatusUp, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("Enrolled minion")

	select {}
}
