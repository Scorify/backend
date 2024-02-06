package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	// GlobalNotification is the key for the global notifications pub/sub in redis
	globalNotification = "global_notification"
)

func (c *Cache) PublishNotification(ctx context.Context, message string) *redis.IntCmd {
	return c.Client.Publish(context.Background(), globalNotification, message)
}

func (c *Cache) SubscribeNotification(ctx context.Context) *redis.PubSub {
	return c.Client.Subscribe(context.Background(), globalNotification)
}
