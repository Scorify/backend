package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/grpc/proto"
)

func (c *MinionClient) SubmitScoreTask(ctx context.Context, statusID uuid.UUID, checkStatus status.Status) (*proto.SubmitScoreTaskResponse, error) {
	return c.client.SubmitScoreTask(ctx, &proto.SubmitScoreTaskRequest{
		MinionId: c.MinionID.String(),
		StatusId: statusID.String(),
		Status: func() proto.Status {
			switch checkStatus {
			case status.StatusUp:
				return proto.Status_up
			case status.StatusDown:
				return proto.Status_down
			default:
				return proto.Status_unknown
			}
		}(),
	})
}
