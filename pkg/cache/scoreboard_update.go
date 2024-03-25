package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/graph/model"
)

const (
	// ScoreboardUpdateChannel is the key for the scoreboard update pub/sub in redis
	ScoreboardUpdateChannel = "scoreboard_update_channel"
)

func PublishScoreboardUpdate(ctx context.Context, redisClient *redis.Client, scoreboardUpdate *model.ScoreboardUpdate) (*redis.IntCmd, error) {
	out, err := json.Marshal(scoreboardUpdate)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, ScoreboardUpdateChannel, out), nil
}

func SubscribeScoreboardUpdate(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, ScoreboardUpdateChannel)
}
