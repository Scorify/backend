package setup

import (
	"bufio"
	"html/template"
	"os"
	"time"

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

func run(cmd *cobra.Command, args []string) {
	choice, err := actionMenu()
	if err != nil {
		logrus.WithError(err).Fatal("failed to show action menu")
	}

	switch choice {
	case actionCreate:
		createMenu()
	case actionUpdate:
		logrus.Fatal("update menu not implemented")
	case actionDelete:
		err = deleteMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show delete menu")
		}
	case actionView:
		err = viewMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show view menu")
		}
	}
}

func createMenu() {
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

	// INTERVAL
	interval, err := promptDuration(
		reader,
		30*time.Second,
		time.Second,
		"Enter the interval of the score task in seconds [30s]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read interval")
	}

	// JWT_TIMEOUT
	jwtTimeout, err := promptDuration(
		reader,
		6*time.Hour,
		time.Hour,
		"Enter the timeout of the JWT (session length) in hours [6h]: ",
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

	// GRPC_HOST
	grpcHost, err := prompt(
		reader,
		"localhost",
		"Enter the host of the gRPC server [localhost]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read gRPC host")
	}

	// GRPC_PORT
	grpcPort, err := promptInt(
		reader,
		50051,
		"Enter the port of the gRPC server [50051]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read gRPC port")
	}

	// GRPC_SECRET
	grpcSecret, err := promptPassword(
		reader,
		"Enter the secret key for the gRPC server [randomly generate]: ",
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read gRPC secret")
	}

	// Write .env file
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
		Interval   time.Duration
		JWTTimeout time.Duration
		JWTSecret  string

		PostgresHost     string
		PostgresPort     int
		PostgresUser     string
		PostgresPassword string
		PostgresDB       string

		RedisHost     string
		RedisPort     int
		RedisPassword string

		GRPCHost   string
		GRPCPort   int
		GRPCSecret string
	}{
		Domain:     domain,
		Port:       port,
		Interval:   interval,
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

		GRPCHost:   grpcHost,
		GRPCPort:   grpcPort,
		GRPCSecret: grpcSecret,
	})
	if err != nil {
		logrus.WithError(err).Fatal("failed to write .env")
	}
}
