package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	// Timeout is the time in hours for which the JWT token is valid.
	Timeout time.Duration = 6 * time.Hour

	// JWTKey is the key used to sign the JWT token.
	JWTKey string = "secret"

	// Domain is the domain of the cookie
	Domain string = "localhost"

	// Port is the port of the server
	Port int = 8080

	// Redis is the configuration for the redis server
	Redis struct {
		Url      string
		Password string
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("failed to load .env file")
	}

	Redis.Url = os.Getenv("REDIS_URL")
	Redis.Password = os.Getenv("REDIS_PASSWORD")
}
