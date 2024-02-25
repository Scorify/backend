package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent/status"
)

var mutationSubmitStatus struct {
	SubmitStatus bool `graphql:"submitStatus(minionKey: $minionKey, status_id: $status_id, status: $status, error: $error)"`
}

// vat test = `mutation {
//   submitStatus(
//     minionKey: $minionKey
//     status_id: $status_id
//     status: $status
//     error: $error
//   )
// }`

func (c *GQLClient) SubmitStatus(ctx context.Context, statusID uuid.UUID, status status.Status, err *string) error {
	return c.httpClient.Mutate(
		ctx,
		&mutationSubmitStatus,
		map[string]interface{}{
			"minionKey": config.Minion.Key,
			"status_id": statusID,
			"status":    status,
			"error":     err,
		},
	)
}
