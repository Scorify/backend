package main

import (
	"github.com/scorify/backend/pkg/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logrus.WithError(err).Fatal("failed to execute command")
	}
}
