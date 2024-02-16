package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	// Domain is the domain of the cookie
	Domain string

	// Port is the port of the server
	Port int

	// JWT is the configuration for the JWT token
	JWT struct {
		// Timeout is the timeout for the JWT token in hours
		Timeout int

		// Key is the secret key for the JWT token
		Secret string
	}

	// Redis is the configuration for the redis server
	Redis struct {
		// Url is the url of the redis server
		Url string

		// Password is the password of the redis server
		Password string
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("failed to load .env file")
	}

	Domain = os.Getenv("DOMAIN")
	if Domain == "" {
		logrus.Fatal("DOMAIN is not set")
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse PORT")
	}

	JWT.Timeout, err = strconv.Atoi(os.Getenv("JWT_TIMEOUT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse JWT_TIMEOUT")
	}

	JWT.Secret = os.Getenv("JWT_SECRET")
	if JWT.Secret == "" {
		logrus.Fatal("JWT_SECRET is not set")
	}

	Redis.Url = os.Getenv("REDIS_URL")
	if Redis.Url == "" {
		logrus.Fatal("REDIS_URL is not set")
	}

	Redis.Password = os.Getenv("REDIS_PASSWORD")
	if Redis.Password == "" {
		logrus.Fatal("REDIS_PASSWORD is not set")
	}
}
