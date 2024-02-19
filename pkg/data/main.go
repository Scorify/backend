package data

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/helpers"
	"github.com/sirupsen/logrus"
)

var (
	Client *ent.Client
	Ctx    context.Context = context.Background()
)

func Init() {
	c, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Postgres.Host,
			config.Postgres.Port,
			config.Postgres.User,
			config.Postgres.Password,
			config.Postgres.DB,
		),
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed opening connection to postgres")
	}

	Client = c

	// Run the auto migration tool.
	if err := c.Schema.Create(Ctx); err != nil {
		logrus.WithError(err).Fatalf("failed creating schema resources")
	}

	exists, err := c.User.Query().
		Where(
			user.UsernameEQ("admin"),
		).Exist(Ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("failed checking if admin user exists")
	}

	if !exists {
		hashedPassword, err := helpers.HashPassword("admin")
		if err != nil {
			logrus.WithError(err).Fatalf("failed hashing admin password")
		}

		_, err = c.User.Create().
			SetUsername("admin").
			SetPassword(hashedPassword).
			SetRole(user.RoleAdmin).
			Save(Ctx)
		if err != nil {
			logrus.WithError(err).Warnf("failed creating admin user")
		}
	}
}
