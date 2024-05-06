package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/auth"
	"github.com/scorify/backend/pkg/cache"
	"github.com/scorify/backend/pkg/checks"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/inject"
	"github.com/scorify/backend/pkg/ent/injectsubmission"
	"github.com/scorify/backend/pkg/ent/predicate"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/scorecache"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/graph/model"
	"github.com/scorify/backend/pkg/helpers"
	"github.com/scorify/backend/pkg/structs"
	"github.com/sirupsen/logrus"
)

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
	out, err := json.Marshal(obj.Config)

	return string(out), err
}

// Configs is the resolver for the configs field.
func (r *checkResolver) Configs(ctx context.Context, obj *ent.Check) ([]*ent.CheckConfig, error) {
	return obj.QueryConfigs().All(ctx)
}

// Statuses is the resolver for the statuses field.
func (r *checkResolver) Statuses(ctx context.Context, obj *ent.Check) ([]*ent.Status, error) {
	return obj.QueryStatuses().All(ctx)
}

// Config is the resolver for the config field.
func (r *checkConfigResolver) Config(ctx context.Context, obj *ent.CheckConfig) (string, error) {
	out, err := json.Marshal(obj.Config)

	return string(out), err
}

// Check is the resolver for the check field.
func (r *checkConfigResolver) Check(ctx context.Context, obj *ent.CheckConfig) (*ent.Check, error) {
	return obj.QueryCheck().Only(ctx)
}

// User is the resolver for the user field.
func (r *checkConfigResolver) User(ctx context.Context, obj *ent.CheckConfig) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// Config is the resolver for the config field.
func (r *configResolver) Config(ctx context.Context, obj *ent.CheckConfig) (string, error) {
	entCheck, err := obj.QueryCheck().Only(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get check: %v", err)
	}

	outConfig := make(map[string]interface{})
	for _, key := range entCheck.EditableFields {
		outConfig[key] = obj.Config[key]
	}

	out, err := json.Marshal(outConfig)

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

// Files is the resolver for the files field.
func (r *injectResolver) Files(ctx context.Context, obj *ent.Inject) ([]string, error) {
	var err error

	files := make([]string, len(obj.Files))

	for i, file := range obj.Files {
		files[i], err = file.APIPath(structs.FileTypeInject, obj.ID)
		if err != nil {
			logrus.Errorf("failed to get file path: %v", err)
		}

	}

	return files, nil
}

// Submissions is the resolver for the submissions field.
func (r *injectResolver) Submissions(ctx context.Context, obj *ent.Inject) ([]*ent.InjectSubmission, error) {
	return obj.QuerySubmissions().All(ctx)
}

// Files is the resolver for the files field.
func (r *injectSubmissionResolver) Files(ctx context.Context, obj *ent.InjectSubmission) ([]string, error) {
	var err error

	files := make([]string, len(obj.Files))

	for i, file := range obj.Files {
		files[i], err = file.APIPath(structs.FileTypeSubmission, obj.ID)
		if err != nil {
			logrus.Errorf("failed to get file path: %v", err)
		}
	}

	return files, nil
}

// User is the resolver for the user field.
func (r *injectSubmissionResolver) User(ctx context.Context, obj *ent.InjectSubmission) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// Inject is the resolver for the inject field.
func (r *injectSubmissionResolver) Inject(ctx context.Context, obj *ent.InjectSubmission) (*ent.Inject, error) {
	return obj.QueryInject().Only(ctx)
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

	token, expiration, err := auth.GenerateJWT(username, entUser.ID, nil)
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

// AdminLogin is the resolver for the adminLogin field.
func (r *mutationResolver) AdminLogin(ctx context.Context, id uuid.UUID) (*model.LoginOutput, error) {
	entUser, err := r.Ent.User.Query().
		Where(
			user.IDEQ(id),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("user id does not exist: %s", id)
	}

	token, expiration, err := auth.GenerateJWT(entUser.Username, entUser.ID, nil)
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

// AdminBecome is the resolver for the adminBecome field.
func (r *mutationResolver) AdminBecome(ctx context.Context, id uuid.UUID) (*model.LoginOutput, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	exists, err := r.Ent.User.Query().
		Where(
			user.IDEQ(id),
		).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("user id does not exist: %s", id)
	}

	if !exists {
		return nil, fmt.Errorf("user id does not exist: %s", id)
	}

	token, expiration, err := auth.GenerateJWT(entUser.Username, entUser.ID, &id)
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
func (r *mutationResolver) CreateCheck(ctx context.Context, name string, source string, weight int, config string, editableFields []string) (*ent.Check, error) {
	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

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

	defaultConfig := make(map[string]interface{})
	defaultEditableFields := []string{}

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

	if editableFields != nil {
		defaultEditableFields = editableFields
	}

	entCheck, err := tx.Check.Create().
		SetName(name).
		SetWeight(weight).
		SetSource(source).
		SetConfig(defaultConfig).
		SetEditableFields(defaultEditableFields).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create check: %v", err)
	}

	entUsers, err := tx.User.Query().
		Where(
			user.RoleEQ(user.RoleUser),
		).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	entCheckConfigs := []*ent.CheckConfigCreate{}

	for _, entUser := range entUsers {
		templateConfig := make(map[string]interface{})
		for key, value := range defaultConfig {
			switch val := value.(type) {
			case string:
				templateConfig[key] = helpers.ConfigTemplate(
					val,
					helpers.Template{
						Number: entUser.Number,
						Name:   entUser.Username,
					},
				)
			default:
				templateConfig[key] = val
			}
		}

		entCheckConfigs = append(entCheckConfigs, tx.CheckConfig.Create().
			SetCheck(entCheck).
			SetUser(entUser).
			SetConfig(templateConfig))
	}

	_, err = tx.CheckConfig.CreateBulk(entCheckConfigs...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create check configs: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
	if err != nil {
		return nil, err
	}

	_, err = cache.PublishScoreboardUpdate(ctx, r.Redis, scoreboard)

	return entCheck, err
}

// UpdateCheck is the resolver for the updateCheck field.
func (r *mutationResolver) UpdateCheck(ctx context.Context, id uuid.UUID, name *string, weight *int, config *string, editableFields []string) (*ent.Check, error) {
	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	entCheck, err := tx.Check.Query().
		Where(
			check.IDEQ(id),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error encounted while getting check: %v", err)
	}

	checkUpdate := tx.Check.UpdateOneID(id)

	if name != nil {
		checkUpdate.SetName(*name)
	}

	if weight != nil {
		checkUpdate.SetWeight(*weight)
	}

	if config != nil || editableFields != nil {
		defaultConfig := make(map[string]interface{})
		defaultEditableFields := entCheck.EditableFields

		for key, value := range entCheck.Config {
			defaultConfig[key] = value
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
		}

		if editableFields != nil {
			defaultEditableFields = editableFields
		}

		checkUpdate.SetConfig(defaultConfig)
		checkUpdate.SetEditableFields(defaultEditableFields)

		checkUpdateResult, err := checkUpdate.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update check: %v", err)
		}

		// generate map of fields and value that were changes
		patchFields := make(map[string]interface{})
		for key, value := range defaultConfig {
			if value != entCheck.Config[key] {
				patchFields[key] = value
			}
		}

		// get all configs to update
		entConfigs, err := tx.CheckConfig.Query().
			WithUser().
			Where(
				checkconfig.HasCheckWith(
					check.IDEQ(id),
				),
			).
			All(ctx)
		if err != nil {
			return nil, err
		}

		// update specific fields that were changed
		for _, entConfig := range entConfigs {
			for key, value := range patchFields {
				switch val := value.(type) {
				case string:
					entConfig.Config[key] = helpers.ConfigTemplate(
						val,
						helpers.Template{
							Number: entConfig.Edges.User.Number,
							Name:   entConfig.Edges.User.Username,
						},
					)
				default:
					entConfig.Config[key] = val
				}
			}

			_, err = entConfig.Update().
				SetConfig(entConfig.Config).
				Save(ctx)
			if err != nil {
				return nil, err
			}
		}

		// update status points if weight is changed
		if weight != nil {
			err = tx.Status.Update().
				Where(
					status.HasCheckWith(
						check.IDEQ(id),
					),
					status.StatusEQ(status.StatusUp),
				).
				SetPoints(*weight).
				Exec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update status: %v", err)
			}

			err = helpers.RecomputeScores(tx, r.Redis, ctx)
			if err != nil {
				return nil, err
			}
		} else {
			scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
			if err != nil {
				return nil, err
			}

			_, err = cache.PublishScoreboardUpdate(ctx, r.Redis, scoreboard)
			if err != nil {
				return nil, err
			}
		}

		err = tx.Commit()
		if err != nil {
			return nil, fmt.Errorf("failed to commit transaction: %v", err)
		}

		return checkUpdateResult, nil
	}

	entCheck, err = checkUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update check: %v", err)
	}

	// update status points if weight is changed
	if weight != nil {
		err = tx.Status.Update().
			Where(
				status.HasCheckWith(
					check.IDEQ(id),
				),
				status.StatusEQ(status.StatusUp),
			).
			SetPoints(*weight).
			Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update status: %v", err)
		}

		err = helpers.RecomputeScores(tx, r.Redis, ctx)
		if err != nil {
			return nil, err
		}
	} else {
		scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
		if err != nil {
			return nil, err
		}

		_, err = cache.PublishScoreboardUpdate(ctx, r.Redis, scoreboard)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return entCheck, nil
}

// DeleteCheck is the resolver for the deleteCheck field.
func (r *mutationResolver) DeleteCheck(ctx context.Context, id uuid.UUID) (bool, error) {
	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	err = tx.Check.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, err
	}

	err = helpers.RecomputeScores(tx, r.Redis, ctx)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	return err == nil, err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, username string, password string, role user.Role, number *int) (*ent.User, error) {
	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return nil, err
	}

	entCreateUser := tx.User.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetRole(role)

	if number != nil {
		entCreateUser.SetNumber(*number)
	}

	entUser, err := entCreateUser.Save(ctx)
	if err != nil {
		return nil, err
	}

	entChecks, err := tx.Check.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	for _, entCheck := range entChecks {
		config := make(map[string]interface{})
		for key, value := range entCheck.Config {
			switch val := value.(type) {
			case string:
				config[key] = helpers.ConfigTemplate(
					val,
					helpers.Template{
						Number: entUser.Number,
						Name:   entUser.Username,
					},
				)
			default:
				config[key] = value
			}
		}

		_, err := tx.CheckConfig.Create().
			SetCheck(entCheck).
			SetConfig(config).
			SetUser(entUser).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
	if err != nil {
		return nil, err
	}

	_, err = cache.PublishScoreboardUpdate(ctx, r.Redis, scoreboard)

	return entUser, err
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id uuid.UUID, username *string, password *string, number *int) (*ent.User, error) {
	userUpdate := r.Ent.User.UpdateOneID(id)

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

	if number != nil {
		userUpdate.SetNumber(*number)
	}

	entUser, err := userUpdate.Save(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
	if err != nil {
		return nil, err
	}

	_, err = cache.PublishScoreboardUpdate(ctx, r.Redis, scoreboard)

	return entUser, err
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return false, fmt.Errorf("invalid user")
	}

	if entUser.ID == id {
		return false, fmt.Errorf("cannot delete yourself")
	}

	err = r.Ent.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, err
	}

	scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
	if err != nil {
		return false, err
	}

	_, err = cache.PublishScoreboardUpdate(ctx, r.Redis, scoreboard)

	return err == nil, err
}

// EditConfig is the resolver for the editConfig field.
func (r *mutationResolver) EditConfig(ctx context.Context, id uuid.UUID, config string) (*ent.CheckConfig, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	entCheckConfig, err := r.Ent.CheckConfig.Query().
		Where(
			checkconfig.IDEQ(id),
			checkconfig.HasUserWith(
				user.IDEQ(
					entUser.ID,
				),
			),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("no check config found")
	}

	newConfig := make(map[string]interface{})

	err = json.Unmarshal([]byte(config), &newConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid config")
	}

	oldConfig := entCheckConfig.Config

	for key, value := range newConfig {
		oldConfig[key] = value
	}

	return r.Ent.CheckConfig.UpdateOneID(id).
		SetConfig(oldConfig).
		Save(ctx)
}

// SendGlobalNotification is the resolver for the sendGlobalNotification field.
func (r *mutationResolver) SendGlobalNotification(ctx context.Context, message string, typeArg model.NotificationType) (bool, error) {
	_, err := cache.PublishNotification(ctx, r.Redis, message, typeArg)
	return err == nil, err
}

// StartEngine is the resolver for the startEngine field.
func (r *mutationResolver) StartEngine(ctx context.Context) (bool, error) {
	err := r.Engine.Start()

	return err == nil, err
}

// StopEngine is the resolver for the stopEngine field.
func (r *mutationResolver) StopEngine(ctx context.Context) (bool, error) {
	err := r.Engine.Stop()

	return err == nil, err
}

// CreateInject is the resolver for the createInject field.
func (r *mutationResolver) CreateInject(ctx context.Context, title string, startTime time.Time, endTime time.Time, files []*graphql.Upload) (*ent.Inject, error) {
	structFiles := make([]structs.File, len(files))

	for i, file := range files {
		structFiles[i] = structs.File{
			ID:   uuid.New(),
			Name: file.Filename,
		}
	}

	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	defer tx.Rollback()

	entInject, err := tx.Inject.Create().
		SetTitle(title).
		SetStartTime(startTime).
		SetEndTime(endTime).
		SetFiles(structFiles).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create inject: %v", err)
	}

	for i, file := range files {
		err = structFiles[i].WriteFile(structs.FileTypeInject, entInject.ID, file.File)
		if err != nil {
			return nil, fmt.Errorf("failed to write file: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return entInject, nil
}

// UpdateInject is the resolver for the updateInject field.
func (r *mutationResolver) UpdateInject(ctx context.Context, id uuid.UUID, title *string, startTime *time.Time, endTime *time.Time) (*ent.Inject, error) {
	if title == nil && startTime == nil && endTime == nil {
		return nil, fmt.Errorf("no fields to update")
	}

	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	defer tx.Rollback()

	injectUpdate := tx.Inject.UpdateOneID(id)

	if title != nil {
		injectUpdate.SetTitle(*title)
	}

	if startTime != nil {
		injectUpdate.SetStartTime(*startTime)
	}

	if endTime != nil {
		injectUpdate.SetEndTime(*endTime)
	}

	entInject, err := injectUpdate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update inject: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return entInject, nil
}

// AddInjectFiles is the resolver for the addInjectFiles field.
func (r *mutationResolver) AddInjectFiles(ctx context.Context, id uuid.UUID, files []*graphql.Upload) (*ent.Inject, error) {
	entInject, err := r.Ent.Inject.Query().
		Where(
			inject.IDEQ(id),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get inject: %v", err)
	}

	structFiles := make([]structs.File, len(files))

	for i, file := range files {
		structFiles[i] = structs.File{
			ID:   uuid.New(),
			Name: file.Filename,
		}

		err := structFiles[i].WriteFile(structs.FileTypeInject, id, file.File)
		if err != nil {
			return nil, fmt.Errorf("failed to write file: %v", err)
		}
	}

	return r.Ent.Inject.UpdateOneID(id).
		SetFiles(append(entInject.Files, structFiles...)).
		Save(ctx)
}

// DeleteInjectFile is the resolver for the deleteInjectFile field.
func (r *mutationResolver) DeleteInjectFile(ctx context.Context, id uuid.UUID, file string) (*ent.Inject, error) {
	panic(fmt.Errorf("not implemented: DeleteInjectFile - deleteInjectFile"))
}

// DeleteInject is the resolver for the deleteInject field.
func (r *mutationResolver) DeleteInject(ctx context.Context, id uuid.UUID) (bool, error) {
	entInject, err := r.Ent.Inject.Query().
		Where(
			inject.IDEQ(id),
		).Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get inject: %v", err)
	}

	tx, err := r.Ent.Tx(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to start transaction: %v", err)
	}

	defer tx.Rollback()

	for _, file := range entInject.Files {
		err = file.DeleteFile(structs.FileTypeInject, id)
		if err != nil {
			return false, fmt.Errorf("failed to delete file: %v", err)
		}
	}

	err = tx.Inject.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to delete inject: %v", err)
	}

	err = tx.Commit()
	return err == nil, err
}

// SubmitInject is the resolver for the submitInject field.
func (r *mutationResolver) SubmitInject(ctx context.Context, injectID uuid.UUID, files []*graphql.Upload) (*ent.InjectSubmission, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	structFiles := make([]structs.File, len(files))

	for i, file := range files {
		structFiles[i] = structs.File{
			ID:   uuid.New(),
			Name: file.Filename,
		}

		err := structFiles[i].WriteFile(structs.FileTypeSubmission, injectID, file.File)
		if err != nil {
			return nil, fmt.Errorf("failed to write file: %v", err)
		}
	}

	return r.Ent.InjectSubmission.Create().
		SetUser(entUser).
		SetInjectID(injectID).
		SetFiles(structFiles).
		Save(ctx)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, nil
	}
	return entUser, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.Ent.User.Query().All(ctx)
}

// Sources is the resolver for the sources field.
func (r *queryResolver) Sources(ctx context.Context) ([]*model.Source, error) {
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
func (r *queryResolver) Check(ctx context.Context, id *uuid.UUID, name *string) (*ent.Check, error) {
	checkQueryPredicates := []predicate.Check{}

	if id != nil {
		checkQueryPredicates = append(checkQueryPredicates, check.IDEQ(*id))
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
func (r *queryResolver) Config(ctx context.Context, id uuid.UUID) (*ent.CheckConfig, error) {
	return r.Ent.CheckConfig.Query().
		Where(
			checkconfig.IDEQ(id),
		).Only(ctx)
}

// Scoreboard is the resolver for the scoreboard field.
func (r *queryResolver) Scoreboard(ctx context.Context, round *int) (*model.Scoreboard, error) {
	scoreboard := &model.Scoreboard{}

	if round == nil {
		if cache.GetObject(ctx, r.Redis, cache.ScoreboardObjectKey, scoreboard) {
			return scoreboard, nil
		}

		scoreboard, err := helpers.Scoreboard(ctx, r.Ent)
		if err != nil {
			return nil, err
		}

		err = cache.SetObject(ctx, r.Redis, cache.ScoreboardObjectKey, scoreboard, 0)
		if err != nil {
			return nil, err
		}

		return scoreboard, nil
	}

	if cache.GetObject(ctx, r.Redis, cache.GetScoreboardObjectKey(*round), scoreboard) {
		return scoreboard, nil
	}

	scoreboard, err := helpers.ScoreboardByRound(ctx, r.Ent, *round)
	if err != nil {
		return nil, err
	}

	err = cache.SetObject(ctx, r.Redis, cache.GetScoreboardObjectKey(*round), scoreboard, 0)
	if err != nil {
		return nil, err
	}

	return scoreboard, nil
}

// Injects is the resolver for the injects field.
func (r *queryResolver) Injects(ctx context.Context) ([]*ent.Inject, error) {
	return r.Ent.Inject.Query().All(ctx)
}

// Inject is the resolver for the inject field.
func (r *queryResolver) Inject(ctx context.Context, id uuid.UUID) (*ent.Inject, error) {
	return r.Ent.Inject.Query().
		Where(
			inject.IDEQ(id),
		).Only(ctx)
}

// InjectSubmissions is the resolver for the injectSubmissions field.
func (r *queryResolver) InjectSubmissions(ctx context.Context) ([]*ent.InjectSubmission, error) {
	return r.Ent.InjectSubmission.Query().All(ctx)
}

// InjectSubmission is the resolver for the injectSubmission field.
func (r *queryResolver) InjectSubmission(ctx context.Context, id uuid.UUID) (*ent.InjectSubmission, error) {
	return r.Ent.InjectSubmission.Query().
		Where(
			injectsubmission.IDEQ(id),
		).Only(ctx)
}

// Statuses is the resolver for the statuses field.
func (r *roundResolver) Statuses(ctx context.Context, obj *ent.Round) ([]*ent.Status, error) {
	return r.Ent.Status.Query().
		Where(
			status.HasRoundWith(
				round.IDEQ(obj.ID),
			),
		).All(ctx)
}

// ScoreCaches is the resolver for the score_caches field.
func (r *roundResolver) ScoreCaches(ctx context.Context, obj *ent.Round) ([]*ent.ScoreCache, error) {
	return r.Ent.ScoreCache.Query().
		Where(
			scorecache.HasRoundWith(
				round.IDEQ(obj.ID),
			),
		).All(ctx)
}

// Round is the resolver for the round field.
func (r *scoreCacheResolver) Round(ctx context.Context, obj *ent.ScoreCache) (*ent.Round, error) {
	return obj.QueryRound().Only(ctx)
}

// User is the resolver for the user field.
func (r *scoreCacheResolver) User(ctx context.Context, obj *ent.ScoreCache) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// Check is the resolver for the check field.
func (r *statusResolver) Check(ctx context.Context, obj *ent.Status) (*ent.Check, error) {
	return obj.QueryCheck().Only(ctx)
}

// Round is the resolver for the round field.
func (r *statusResolver) Round(ctx context.Context, obj *ent.Status) (*ent.Round, error) {
	return obj.QueryRound().Only(ctx)
}

// User is the resolver for the user field.
func (r *statusResolver) User(ctx context.Context, obj *ent.Status) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// GlobalNotification is the resolver for the globalNotification field.
func (r *subscriptionResolver) GlobalNotification(ctx context.Context) (<-chan *model.Notification, error) {
	notification_chan := make(chan *model.Notification, 1)

	go func() {
		sub := cache.SubscribeNotification(ctx, r.Redis)

		ch := sub.Channel()
		for {
			select {
			case msg := <-ch:
				notification := model.Notification{}
				err := json.Unmarshal([]byte(msg.Payload), &notification)
				if err != nil {
					logrus.WithError(err).Error("failed to unmarshal notification")
					continue
				}

				notification_chan <- &notification
			case <-ctx.Done():
				close(notification_chan)
				sub.Close()
				return
			}
		}
	}()

	return notification_chan, nil
}

// EngineState is the resolver for the engineState field.
func (r *subscriptionResolver) EngineState(ctx context.Context) (<-chan model.EngineState, error) {
	engineStateChan := make(chan model.EngineState, 1)

	go func() {
		sub := cache.SubscribeEngineState(ctx, r.Redis)
		defer sub.Close()
		defer close(engineStateChan)

		ch := sub.Channel()

		state, err := r.Engine.State()
		if err != nil {
			logrus.WithError(err).Error("failed to get engine state")
			return
		}

		engineStateChan <- state

		for {
			select {
			case msg := <-ch:
				state := model.EngineState(msg.Payload)
				engineStateChan <- state
			case <-ctx.Done():
				return
			}
		}
	}()

	return engineStateChan, nil
}

// ScoreboardUpdate is the resolver for the scoreboardUpdate field.
func (r *subscriptionResolver) ScoreboardUpdate(ctx context.Context) (<-chan *model.Scoreboard, error) {
	scoreboardUpdateChan := make(chan *model.Scoreboard, 1)

	go func() {
		scoreboardSub := cache.SubscribeScoreboardUpdate(ctx, r.Redis)
		scoreboardChan := scoreboardSub.Channel()

		for {
			select {
			case msg := <-scoreboardChan:
				scoreboardUpdate := &model.Scoreboard{}
				err := json.Unmarshal([]byte(msg.Payload), scoreboardUpdate)
				if err != nil {
					logrus.WithError(err).Error("failed to unmarshal round update")
					continue
				}

				scoreboardUpdateChan <- scoreboardUpdate
			case <-ctx.Done():
				close(scoreboardUpdateChan)
				scoreboardSub.Close()
				return
			}
		}
	}()

	return scoreboardUpdateChan, nil
}

// LatestRound is the resolver for the latestRound field.
func (r *subscriptionResolver) LatestRound(ctx context.Context) (<-chan *ent.Round, error) {
	latestRoundChan := make(chan *ent.Round, 1)

	go func() {
		latestRound := &ent.Round{}
		if cache.GetObject(ctx, r.Redis, cache.LatestRoundObjectKey, latestRound) {
			latestRoundChan <- latestRound
		}

		latestRoundSub := cache.SubscribeLatestRound(ctx, r.Redis)
		latestRoundSubChan := latestRoundSub.Channel()

		for {
			select {
			case msg := <-latestRoundSubChan:
				latestRound := &ent.Round{}
				err := json.Unmarshal([]byte(msg.Payload), latestRound)
				if err != nil {
					logrus.WithError(err).Error("failed to unmarshal latest round")
					continue
				}

				latestRoundChan <- latestRound
			case <-ctx.Done():
				close(latestRoundChan)
				latestRoundSub.Close()
				return
			}
		}
	}()

	return latestRoundChan, nil
}

// Configs is the resolver for the configs field.
func (r *userResolver) Configs(ctx context.Context, obj *ent.User) ([]*ent.CheckConfig, error) {
	return obj.QueryConfigs().All(ctx)
}

// Statuses is the resolver for the statuses field.
func (r *userResolver) Statuses(ctx context.Context, obj *ent.User) ([]*ent.Status, error) {
	return obj.QueryStatuses().All(ctx)
}

// ScoreCaches is the resolver for the score_caches field.
func (r *userResolver) ScoreCaches(ctx context.Context, obj *ent.User) ([]*ent.ScoreCache, error) {
	return obj.QueryScoreCaches().All(ctx)
}

// InjectSubmissions is the resolver for the inject_submissions field.
func (r *userResolver) InjectSubmissions(ctx context.Context, obj *ent.User) ([]*ent.InjectSubmission, error) {
	return obj.QuerySubmissions().All(ctx)
}

// Check returns CheckResolver implementation.
func (r *Resolver) Check() CheckResolver { return &checkResolver{r} }

// CheckConfig returns CheckConfigResolver implementation.
func (r *Resolver) CheckConfig() CheckConfigResolver { return &checkConfigResolver{r} }

// Config returns ConfigResolver implementation.
func (r *Resolver) Config() ConfigResolver { return &configResolver{r} }

// Inject returns InjectResolver implementation.
func (r *Resolver) Inject() InjectResolver { return &injectResolver{r} }

// InjectSubmission returns InjectSubmissionResolver implementation.
func (r *Resolver) InjectSubmission() InjectSubmissionResolver { return &injectSubmissionResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Round returns RoundResolver implementation.
func (r *Resolver) Round() RoundResolver { return &roundResolver{r} }

// ScoreCache returns ScoreCacheResolver implementation.
func (r *Resolver) ScoreCache() ScoreCacheResolver { return &scoreCacheResolver{r} }

// Status returns StatusResolver implementation.
func (r *Resolver) Status() StatusResolver { return &statusResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type checkResolver struct{ *Resolver }
type checkConfigResolver struct{ *Resolver }
type configResolver struct{ *Resolver }
type injectResolver struct{ *Resolver }
type injectSubmissionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roundResolver struct{ *Resolver }
type scoreCacheResolver struct{ *Resolver }
type statusResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
