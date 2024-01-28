package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/checks"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/graph/model"
	"github.com/scorify/backend/pkg/helpers"
)

// ID is the resolver for the id field.
func (r *checkResolver) ID(ctx context.Context, obj *ent.Check) (string, error) {
	return obj.ID.String(), nil
}

// Source is the resolver for the source field.
func (r *checkResolver) Source(ctx context.Context, obj *ent.Check) (*model.Source, error) {
	schema, ok := checks.Checks[obj.Source]
	if !ok {
		return nil, fmt.Errorf("source \"%s\" does not exist", obj.Source)
	}

	return &model.Source{
		Name:   obj.Source,
		Schema: schema.Schema,
	}, nil
}

// Config is the resolver for the config field.
func (r *checkResolver) Config(ctx context.Context, obj *ent.Check) (string, error) {
	out, err := json.Marshal(obj.DefaultConfig)

	return string(out), err
}

// ID is the resolver for the id field.
func (r *configResolver) ID(ctx context.Context, obj *ent.CheckConfig) (string, error) {
	return obj.ID.String(), nil
}

// Config is the resolver for the config field.
func (r *configResolver) Config(ctx context.Context, obj *ent.CheckConfig) (string, error) {
	out, err := json.Marshal(obj.Config)

	return string(out), err
}

// Check is the resolver for the check field.
func (r *configResolver) Check(ctx context.Context, obj *ent.CheckConfig) (*ent.Check, error) {
	return obj.QueryCheck().Only(ctx)
}

// User is the resolver for the user field.
func (r *configResolver) User(ctx context.Context, obj *ent.CheckConfig) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
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

// CreateCheck is the resolver for the createCheck field.
func (r *mutationResolver) CreateCheck(ctx context.Context, name string, source string, config string) (*ent.Check, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return nil, fmt.Errorf("invalid permissions")
	}

	configSchema, ok := checks.Checks[source]
	if !ok {
		return nil, fmt.Errorf("source \"%s\" does not exist", source)
	}

	var schemaMap map[string]interface{}
	err = json.Unmarshal([]byte(configSchema.Schema), &schemaMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal schema: %v", err)
	}

	var configMap map[string]interface{}
	err = json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	defaultConfig := map[string]interface{}{}
	for key, value := range schemaMap {
		switch value {
		case "string":
			configValue, ok := configMap[key]
			if !ok {
				return nil, fmt.Errorf("invalid config, missing key \"%s\"", key)
			}

			configString, ok := configValue.(string)
			if !ok {
				return nil, fmt.Errorf("invalid config, key \"%s\" is not a string", key)
			}

			defaultConfig[key] = configString
		case "int":
			configValue, ok := configMap[key]
			if !ok {
				return nil, fmt.Errorf("invalid config, missing key \"%s\"", key)
			}

			configFloat, ok := configValue.(float64)
			if !ok {
				return nil, fmt.Errorf("invalid config, key \"%s\" is not an int", key)
			}

			defaultConfig[key] = int(configFloat)
		case "bool":
			configValue, ok := configMap[key]
			if !ok {
				return nil, fmt.Errorf("invalid config, missing key \"%s\"", key)
			}

			configBool, ok := configValue.(bool)
			if !ok {
				return nil, fmt.Errorf("invalid config, key \"%s\" is not a boolean", key)
			}

			defaultConfig[key] = configBool
		default:
			return nil, fmt.Errorf("invalid schema, unknown type \"%s\" for key \"%s\"", value, key)
		}
	}

	_, ok = checks.Checks[source]
	if !ok {
		return nil, fmt.Errorf("source \"%s\" does not exist", source)
	}

	entCheck, err := r.Ent.Check.Create().
		SetName(name).
		SetSource(source).
		SetDefaultConfig(defaultConfig).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create check: %v", err)
	}

	entUsers, err := r.Ent.User.Query().
		Where(
			user.RoleEQ(user.RoleUser),
		).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	entCheckConfigs := []*ent.CheckConfigCreate{}

	for _, entUser := range entUsers {
		entCheckConfigs = append(entCheckConfigs, r.Ent.CheckConfig.Create().
			SetCheck(entCheck).
			SetUser(entUser).
			SetConfig(defaultConfig))
	}

	_, err = r.Ent.CheckConfig.CreateBulk(entCheckConfigs...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create check configs: %v", err)
	}

	return entCheck, nil
}

// UpdateCheck is the resolver for the updateCheck field.
func (r *mutationResolver) UpdateCheck(ctx context.Context, id string, name *string, config *string) (*ent.Check, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return nil, fmt.Errorf("invalid permissions")
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("encounter error while parsing id: %v", err)
	}

	entCheck, err := r.Ent.Check.Query().
		Where(
			check.IDEQ(uuid),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error encounted while getting check: %v", err)
	}

	checkUpdate := r.Ent.Check.UpdateOneID(uuid)

	if name != nil {
		checkUpdate.SetName(*name)
	}

	if config != nil {

		configSchema, ok := checks.Checks[entCheck.Source]
		if !ok {
			return nil, fmt.Errorf("source \"%s\" does not exist", entCheck.Source)
		}

		var schemaMap map[string]interface{}
		err = json.Unmarshal([]byte(configSchema.Schema), &schemaMap)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal schema: %v", err)
		}

		var configMap map[string]interface{}
		err = json.Unmarshal([]byte(*config), &configMap)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal config: %v", err)
		}

		defaultConfig := map[string]interface{}{}
		for key, value := range schemaMap {
			switch value {
			case "string":
				configValue, ok := configMap[key]
				if !ok {
					return nil, fmt.Errorf("invalid config, missing key \"%s\"", key)
				}

				configString, ok := configValue.(string)
				if !ok {
					return nil, fmt.Errorf("invalid config, key \"%s\" is not a string", key)
				}

				defaultConfig[key] = configString
			case "int":
				configValue, ok := configMap[key]
				if !ok {
					return nil, fmt.Errorf("invalid config, missing key \"%s\"", key)
				}

				configFloat, ok := configValue.(float64)
				if !ok {
					return nil, fmt.Errorf("invalid config, key \"%s\" is not an int", key)
				}

				defaultConfig[key] = int(configFloat)
			case "bool":
				configValue, ok := configMap[key]
				if !ok {
					return nil, fmt.Errorf("invalid config, missing key \"%s\"", key)
				}

				configBool, ok := configValue.(bool)
				if !ok {
					return nil, fmt.Errorf("invalid config, key \"%s\" is not a boolean", key)
				}

				defaultConfig[key] = configBool
			default:
				return nil, fmt.Errorf("invalid schema, unknown type \"%s\" for key \"%s\"", value, key)
			}
		}

		checkUpdate.SetDefaultConfig(defaultConfig)

		checkUpdateResult, err := checkUpdate.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update check: %v", err)
		}

		_, err = r.Ent.CheckConfig.Update().
			Where(
				checkconfig.HasCheckWith(
					check.IDEQ(uuid),
				),
			).
			SetConfig(defaultConfig).
			Save(ctx)

		return checkUpdateResult, err
	}

	return checkUpdate.Save(ctx)
}

// DeleteCheck is the resolver for the deleteCheck field.
func (r *mutationResolver) DeleteCheck(ctx context.Context, id string) (bool, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return false, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return false, fmt.Errorf("invalid permissions")
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return false, fmt.Errorf("encounter error while parsing id: %v", err)
	}

	err = r.Ent.Check.DeleteOneID(uuid).Exec(ctx)

	return err == nil, err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, username string, password string, role user.Role, number *int) (*ent.User, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return nil, fmt.Errorf("invalid permissions")
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return nil, err
	}

	entCreateUser := r.Ent.User.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetRole(role)

	if number != nil {
		entCreateUser.SetNumber(*number)
	}

	return entCreateUser.Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, username *string, password *string, role *user.Role, number *int) (*ent.User, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return nil, fmt.Errorf("invalid permissions")
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("encounter error while parsing id: %v", err)
	}

	userUpdate := r.Ent.User.UpdateOneID(uuid)

	if username != nil {
		userUpdate.SetUsername(*username)
	}

	if password != nil {
		hashedPassword, err := helpers.HashPassword(*password)
		if err != nil {
			return nil, err
		}

		userUpdate.SetPassword(hashedPassword)
	}

	if role != nil {
		userUpdate.SetRole(*role)
	}

	if number != nil {
		userUpdate.SetNumber(*number)
	}

	return userUpdate.Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return false, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return false, fmt.Errorf("invalid permissions")
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return false, fmt.Errorf("encounter error while parsing id: %v", err)
	}

	err = r.Ent.User.DeleteOneID(uuid).Exec(ctx)

	return err == nil, err
}

// EditConfig is the resolver for the editConfig field.
func (r *mutationResolver) EditConfig(ctx context.Context, id string, config string) (*ent.CheckConfig, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("encounter error while parsing id: %v", err)
	}

	entCheckConfig, err := r.Ent.CheckConfig.Query().
		Where(
			checkconfig.IDEQ(uuid),
			checkconfig.HasUserWith(
				user.IDEQ(
					entUser.ID,
				),
			),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("no check config found")
	}

	var newConfig map[string]interface{}
	err = json.Unmarshal([]byte(config), &newConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid config")
	}

	oldConfig := entCheckConfig.Config

	for key, value := range newConfig {
		oldConfig[key] = value
	}

	return r.Ent.CheckConfig.UpdateOneID(uuid).Save(ctx)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	return auth.Parse(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Sources is the resolver for the sources field.
func (r *queryResolver) Sources(ctx context.Context) ([]*model.Source, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if entUser.Role != user.RoleAdmin {
		return nil, fmt.Errorf("invalid permissions")
	}

	var checkSources []*model.Source

	for name, schema := range checks.Checks {
		checkSources = append(checkSources, &model.Source{
			Name:   name,
			Schema: schema.Schema,
		})
	}

	return checkSources, nil
}

// Source is the resolver for the source field.
func (r *queryResolver) Source(ctx context.Context, name string) (*model.Source, error) {
	checkSource, ok := checks.Checks[name]
	if !ok {
		return nil, fmt.Errorf("source \"%s\" does not exist", name)
	}

	return &model.Source{
		Name:   name,
		Schema: checkSource.Schema,
	}, nil
}

// Checks is the resolver for the checks field.
func (r *queryResolver) Checks(ctx context.Context) ([]*ent.Check, error) {
	return r.Ent.Check.Query().All(ctx)
}

// Check is the resolver for the check field.
func (r *queryResolver) Check(ctx context.Context, id *string, name *string) (*ent.Check, error) {
	checkQueryPredicates := []predicate.Check{}

	if id != nil {
		uuid, err := uuid.Parse(*id)
		if err != nil {
			return nil, fmt.Errorf("encounter error while parsing id: %v", err)
		}

		checkQueryPredicates = append(checkQueryPredicates, check.IDEQ(uuid))
	}

	if name != nil {
		checkQueryPredicates = append(checkQueryPredicates, check.NameEQ(*name))
	}

	return r.Ent.Check.Query().
		Where(
			checkQueryPredicates...,
		).Only(ctx)
}

// Configs is the resolver for the configs field.
func (r *queryResolver) Configs(ctx context.Context) ([]*ent.CheckConfig, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if entUser.Role == user.RoleAdmin {
		return r.Ent.CheckConfig.Query().All(ctx)
	} else {
		return r.Ent.CheckConfig.Query().
			Where(
				checkconfig.HasUserWith(
					user.IDEQ(
						entUser.ID,
					),
				),
			).All(ctx)
	}
}

// Config is the resolver for the config field.
func (r *queryResolver) Config(ctx context.Context, id string) (*ent.CheckConfig, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("encounter error while parsing id: %v", err)
	}

	return r.Ent.CheckConfig.Query().
		Where(
			checkconfig.IDEQ(uuid),
		).Only(ctx)
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Check returns CheckResolver implementation.
func (r *Resolver) Check() CheckResolver { return &checkResolver{r} }

// Config returns ConfigResolver implementation.
func (r *Resolver) Config() ConfigResolver { return &configResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type checkResolver struct{ *Resolver }
type configResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
