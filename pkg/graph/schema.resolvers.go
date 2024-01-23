package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"fmt"

	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/graph/model"
	"github.com/scorify/backend/pkg/helpers"
)

// ID is the resolver for the id field.
func (r *checkResolver) ID(ctx context.Context, obj *ent.Check) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Source is the resolver for the source field.
func (r *checkResolver) Source(ctx context.Context, obj *ent.Check) (*model.CheckSource, error) {
	panic(fmt.Errorf("not implemented: Source - source"))
}

// ID is the resolver for the id field.
func (r *checkConfigResolver) ID(ctx context.Context, obj *ent.CheckConfig) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Name is the resolver for the name field.
func (r *checkConfigResolver) Name(ctx context.Context, obj *ent.CheckConfig) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Check is the resolver for the check field.
func (r *checkConfigResolver) Check(ctx context.Context, obj *ent.CheckConfig) (*ent.Check, error) {
	panic(fmt.Errorf("not implemented: Check - check"))
}

// User is the resolver for the user field.
func (r *checkConfigResolver) User(ctx context.Context, obj *ent.CheckConfig) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.LoginOutput, error) {
	entUser, err := r.Ent.User.Query().
		Where(
			user.UsernameEQ(username),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	success := helpers.ComparePasswords(entUser.Password, password)
	if !success {
		return nil, fmt.Errorf("invalid username or password")
	}

	token, expiration, err := auth.GenerateJWT(username, string(entUser.Role))
	if err != nil {
		return nil, err
	}

	return &model.LoginOutput{
		Name:     "auth",
		Token:    token,
		Expires:  expiration,
		Path:     "/",
		Domain:   config.Domain,
		Secure:   false,
		HTTPOnly: false,
	}, nil
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string) (bool, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return false, fmt.Errorf("invalid user")
	}

	success := helpers.ComparePasswords(entUser.Password, oldPassword)
	if !success {
		return false, fmt.Errorf("invalid old password")
	}

	hashedPassword, err := helpers.HashPassword(newPassword)
	if err != nil {
		return false, err
	}

	_, err = r.Ent.User.UpdateOneID(entUser.ID).
		SetPassword(hashedPassword).
		Save(ctx)
	return err == nil, err
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	return auth.Parse(ctx)
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Check returns CheckResolver implementation.
func (r *Resolver) Check() CheckResolver { return &checkResolver{r} }

// CheckConfig returns CheckConfigResolver implementation.
func (r *Resolver) CheckConfig() CheckConfigResolver { return &checkConfigResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type checkResolver struct{ *Resolver }
type checkConfigResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
