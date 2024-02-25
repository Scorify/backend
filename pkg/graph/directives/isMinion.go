package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/scorify/backend/pkg/auth"
)

func IsMinion(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	isMinion := auth.IsMinion(ctx)
	if !isMinion {
		return nil, fmt.Errorf("missing minion authentication token")
	}

	return next(ctx)
}
