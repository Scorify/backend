package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/ent/user"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*user.Role) (res interface{}, err error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		if *role == entUser.Role {
			return next(ctx)
		}
	}

	return nil, fmt.Errorf("user does any of the required roles: %v", roles)
}
