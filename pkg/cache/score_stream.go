package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/ent"
)

const (
	// ScoreStreamChannel is the key for the score stream pub/sub in redis
	scoreStreamChannel = "score_stream_channel"
)

func PublishScoreStream(ctx context.Context, redisClient *redis.Client, status *ent.Status) (*redis.IntCmd, error) {
	out, err := json.Marshal(status)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, scoreStreamChannel, out), nil
}

func SubscribeScoreStream(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, scoreStreamChannel)
}
