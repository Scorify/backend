package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/scorify/backend/pkg/auth"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	_, err = auth.Parse(ctx)
	if err != nil {
		return nil, err
	}

	return next(ctx)
}
