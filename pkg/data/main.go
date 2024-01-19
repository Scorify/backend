package data

import (
	"context"

	_ "github.com/mattn/go-sqlite3"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/helpers"
	"github.com/sirupsen/logrus"
)

var (
	Client *ent.Client
	Ctx    context.Context = context.Background()
)

func init() {
	c, err := ent.Open("sqlite3", "file:database.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		logrus.WithError(err).Fatalf("failed opening connection to sqlite")
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
