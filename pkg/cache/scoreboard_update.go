package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/graph/model"
)

const (
	// ScoreboardRoundUpdateChannel is the channel to publish scoreboard round updates
	ScoreboardRoundUpdateChannel = "scoreboard_round_update_channel"

	// ScoreboardStatusUpdateChannel is the channel to publish scoreboard status updates
	ScoreboardStatusUpdateChannel = "scoreboard_status_update_channel"

	// ScoreboardScoreUpdateChannel is the channel to publish scoreboard score updates
	ScoreboardScoreUpdateChannel = "scoreboard_score_update_channel"
)

func PublishScoreboardRoundUpdate(ctx context.Context, redisClient *redis.Client, roundUpdate *model.RoundUpdateScoreboard) (*redis.IntCmd, error) {
	out, err := json.Marshal(roundUpdate)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, ScoreboardRoundUpdateChannel, out), nil
}

func SubscribeScoreboardRoundUpdate(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, ScoreboardRoundUpdateChannel)
}

func PublishScoreboardStatusUpdate(ctx context.Context, redisClient *redis.Client, statusUpdate *model.StatusUpdateScoreboard) (*redis.IntCmd, error) {
	out, err := json.Marshal(statusUpdate)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, ScoreboardStatusUpdateChannel, out), nil
}

func SubscribeScoreboardStatusUpdate(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, ScoreboardStatusUpdateChannel)
}

func PublishScoreboardScoreUpdate(ctx context.Context, redisClient *redis.Client, scoreUpdate *model.ScoreUpdateScoreboard) (*redis.IntCmd, error) {
	out, err := json.Marshal(scoreUpdate)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, ScoreboardScoreUpdateChannel, out), nil
}

func SubscribeScoreboardScoreUpdate(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, ScoreboardScoreUpdateChannel)
}
