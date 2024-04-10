package helpers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/status"
)

func RecomputeScores(tx *ent.Tx, ctx context.Context) error {
	entRounds, err := tx.Round.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("failed to get rounds: %v", err)
	}

	_, err = tx.ScoreCache.Delete().Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete old score caches: %v", err)
	}

	scoreCacheUpdates := make([]*ent.ScoreCacheCreate, 0)

	type ScoreCaches_t struct {
		UserID uuid.UUID `json:"user_id"`
		Score  int       `json:"sum"`
	}

	for _, entRound := range entRounds {
		ScoreCaches := make([]ScoreCaches_t, 0)

		err = tx.Status.Query().
			Where(
				status.RoundID(entRound.ID),
			).
			GroupBy(
				status.FieldUserID,
			).
			Aggregate(
				ent.Sum(
					status.FieldPoints,
				),
			).
			Scan(ctx, &ScoreCaches)
		if err != nil {
			return fmt.Errorf("failed to get scores: %v", err)
		}

		for _, ScoreCache := range ScoreCaches {
			scoreCacheUpdates = append(
				scoreCacheUpdates,
				tx.ScoreCache.Create().
					SetUserID(ScoreCache.UserID).
					SetRoundID(entRound.ID).
					SetPoints(ScoreCache.Score),
			)
		}
	}

	_, err = tx.ScoreCache.CreateBulk(scoreCacheUpdates...).Save(ctx)

	return err
}
