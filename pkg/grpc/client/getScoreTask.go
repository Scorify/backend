package client

import (
	"context"

	"github.com/scorify/backend/pkg/grpc/proto"
)

func (c *MinionClient) GetScoreTask(ctx context.Context) (*proto.GetScoreTaskResponse, error) {
	return c.client.GetScoreTask(ctx, &proto.GetScoreTaskRequest{
		MinionId: c.MinionID.String(),
	})
}
