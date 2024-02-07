package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/graph/model"
)

const (
	// GlobalNotification is the key for the global notifications pub/sub in redis
	globalNotification = "global_notification"
)

func (c *Cache) PublishNotification(ctx context.Context, message string, notification_type model.NotificationType) (*redis.IntCmd, error) {
	out, err := json.Marshal(model.Notification{
		Message: message,
		Type:    notification_type,
	})
	if err != nil {
		return nil, err
	}

	return c.Client.Publish(context.Background(), globalNotification, out), nil
}

func (c *Cache) SubscribeNotification(ctx context.Context) *redis.PubSub {
	return c.Client.Subscribe(context.Background(), globalNotification)
}
