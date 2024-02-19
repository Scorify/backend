package setup

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "setup",
	Short:   "Setup configuration for the server",
	Long:    "Setup configuration for the server",
	Aliases: []string{"init", "i"},

	Run: run,
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword() (string, error) {
	bytes := make([]byte, 64)

	for i := 0; i < 64; i++ {
		index := rand.Intn(len(charset))
		bytes[i] = charset[index]
	}

	return string(bytes), nil
}

func prompt(reader *bufio.Reader, defaultValue string, message string) (string, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if text == "\n" {
		return defaultValue, nil
	}

	return strings.TrimSpace(text), nil
}

func promptPassword(reader *bufio.Reader, message string) (string, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if text == "\n" {
		password, err := generatePassword()
		if err != nil {
			return "", err
		}

		return password, nil
	}

	return strings.TrimSpace(text), nil
}

func promptInt(reader *bufio.Reader, defaultValue int, message string) (int, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	if text == "\n" {
		return defaultValue, nil
	}

	return strconv.Atoi(strings.TrimSpace(text))
}

func run(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	// DOMAIN
	domain, err := prompt(
		reader,
		"localhost",
		"Enter the domain of the server [localhost]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read domain")
	}

	// PORT
	port, err := promptInt(
		reader,
		8080,
		"Enter the port of the server [8080]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read port")
	}

	// JWT_TIMEOUT
	jwtTimeout, err := promptInt(
		reader,
		60,
		"Enter the timeout of the JWT (session length) in hours [6]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read JWT timeout")
	}

	// JWT_SECRET
	jwtSecret, err := promptPassword(
		reader,
		"Enter the secret key for the JWT token [randomly generate]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read JWT secret")
	}

	// POSTGRES_HOST
	postgresHost, err := prompt(
		reader,
		"postgres",
		"Enter the host of the postgres database [postgres]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read postgres host")
	}

	// POSTGRES_PORT
	postgresPort, err := promptInt(
		reader,
		5432,
		"Enter the port of the postgres database [5432]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read postgres port")
	}

	// POSTGRES_USER
	postgresUser, err := prompt(
		reader,
		"scorify",
		"Enter the user of the postgres database [scorify]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read postgres user")
	}

	// POSTGRES_PASSWORD
	postgresPassword, err := promptPassword(
		reader,
		"Enter the password of the postgres database [randomly generate]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read postgres password")
	}

	// POSTGRES_DB
	postgresDB, err := prompt(
		reader,
		"scorify",
		"Enter the name of the postgres database [scorify]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read postgres database")
	}

	// REDIS_HOST
	redisHost, err := prompt(
		reader,
		"redis",
		"Enter the host of the redis server [redis]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read redis host")
	}

	// REDIS_PORT
	redisPort, err := promptInt(
		reader,
		6379,
		"Enter the port of the redis server [6379]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read redis port")
	}

	// REDIS_PASSWORD
	redisPassword, err := promptPassword(
		reader,
		"Enter the password of the redis server [randomly generate]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read redis password")
	}

	envTmpl, err := os.ReadFile(".env.tmpl")
	if err != nil {
		logrus.WithError(err).Fatal("failed to read .env.tmpl")
	}

	tmpl, err := template.New("env").Parse(string(envTmpl))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse .env.tmpl")
	}

	envFile, err := os.Create(".env")
	if err != nil {
		logrus.WithError(err).Fatal("failed to create .env")
	}

	err = tmpl.Execute(envFile, struct {
		Domain     string
		Port       int
		JWTTimeout int
		JWTSecret  string

		PostgresHost     string
		PostgresPort     int
		PostgresUser     string
		PostgresPassword string
		PostgresDB       string

		RedisHost     string
		RedisPort     int
		RedisPassword string
	}{
		Domain:     domain,
		Port:       port,
		JWTTimeout: jwtTimeout,
		JWTSecret:  jwtSecret,

		PostgresHost:     postgresHost,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDB:       postgresDB,

		RedisHost:     redisHost,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,
	})
	if err != nil {
		logrus.WithError(err).Fatal("failed to write .env")
	}
}