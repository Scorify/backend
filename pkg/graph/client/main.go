package client

import (
	"fmt"

	"github.com/hasura/go-graphql-client"
	"github.com/scorify/backend/pkg/config"
)

type GQLClient struct {
	httpClient *graphql.Client
	wsClient   *graphql.SubscriptionClient
}

func NewClient() (*GQLClient, error) {
	httpGqlEndpoint := fmt.Sprintf("http://%s:%d/query", config.Domain, config.Port)
	// wsGqlEndpoint := fmt.Sprintf("ws://%s:%d/query", config.Domain, config.Port)

	return &GQLClient{
		httpClient: graphql.NewClient(
			httpGqlEndpoint,
			nil,
		),
		wsClient: graphql.NewSubscriptionClient(
			"ws://localhost:8080/query",
		),
	}, nil
}
