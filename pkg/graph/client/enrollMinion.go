package client

import (
	"context"
	"fmt"

	"github.com/scorify/backend/pkg/config"
	"github.com/sirupsen/logrus"
)

// var subscriptionEnrollMinion struct {
// 	EnrollMinion struct {
// 		ID     string `json:"id"`
// 		Config struct {
// 			Config map[string]interface{} `json:"config"`
// 		} `json:"config"`
// 		Check struct {
// 			Source struct {
// 				Name string `json:"name"`
// 			} `json:"source"`
// 		} `json:"check"`
// 	}
// }
// var test = `
// subscription($minionKey: String!) {
// 	enrollMinion(minionKey: $minionKey) {
// 		id
// 		check {
// 			source {
// 				name
// 			}
// 		}
// 		config {
// 			config
// 		}
// 	}
// }`

var subscription struct {
	EnrollMinion struct {
		ID    string `json:"id"`
		Check struct {
			Source struct {
				Name string `json:"name"`
			} `json:"source"`
		} `json:"check"`
		Config struct {
			Config [][2]interface{} `json:"config"`
		} `json:"config"`
	} `json:"enrollMinion(minionKey: $minionKey)"`
}

func (c *GQLClient) EnrollMinion(ctx context.Context) (string, error) {
	return c.wsClient.Subscribe(
		&subscription,
		map[string]interface{}{
			"minionKey": config.Minion.Key,
		},
		func(dataBytes []byte, errValue error) error {
			if errValue != nil {
				logrus.WithError(errValue).Info("error while subscribing to enrollMinion")
			}

			fmt.Println(string(dataBytes))
			return nil
		},
	)
}
