package helpers

import (
	"context"

	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/scorecache"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/graph/model"
)

func Scoreboard(ctx context.Context, entClient *ent.Client) (*model.Scoreboard, error) {
	scoreboard := &model.Scoreboard{}

	entUsers, err := entClient.User.Query().
		Where(
			user.RoleEQ(user.RoleUser),
		).
		Order(
			ent.Asc(user.FieldNumber),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Teams = entUsers

	entChecks, err := entClient.Check.Query().
		Order(
			ent.Asc(check.FieldName),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Checks = entChecks

	entRound, err := entClient.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Round = entRound

	entStatuses, err := entClient.Status.Query().
		Where(
			status.HasRoundWith(
				round.IDEQ(entRound.ID),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	lookup := make(map[uuid.UUID]int)
	for i, entUser := range entUsers {
		lookup[entUser.ID] = i
	}
	for i, entCheck := range entChecks {
		lookup[entCheck.ID] = i
	}

	scoreboard.Statuses = make([][]*ent.Status, len(entChecks))
	for i := range scoreboard.Statuses {
		scoreboard.Statuses[i] = make([]*ent.Status, len(entUsers))
	}

	for _, entStatus := range entStatuses {
		check_index, ok := lookup[entStatus.CheckID]
		if !ok {
			continue
		}

		user_index, ok := lookup[entStatus.UserID]
		if !ok {
			continue
		}

		scoreboard.Statuses[check_index][user_index] = entStatus
	}

	var TeamScore []struct {
		TeamID uuid.UUID `json:"user_id"`
		Sum    int       `json:"sum"`
	}

	err = entClient.ScoreCache.Query().
		GroupBy(
			scorecache.FieldUserID,
		).
		Aggregate(
			ent.Sum(
				scorecache.FieldPoints,
			),
		).
		Scan(ctx, &TeamScore)
	if err != nil {
		return nil, err
	}

	scoreboard.Scores = make([]int, len(entUsers))
	for _, teamScore := range TeamScore {
		user_index, ok := lookup[teamScore.TeamID]
		if !ok {
			continue
		}

		scoreboard.Scores[user_index] = teamScore.Sum
	}

	return scoreboard, nil
}
