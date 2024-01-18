package config

import "time"

var (
	// Timeout is the time in hours for which the JWT token is valid.
	Timeout time.Duration = 6 * time.Hour

	// JWTKey is the key used to sign the JWT token.
	JWTKey string = "secret"

	// Domain is the domain of the cookie
	Domain string = "localhost"

	// Port is the port of the server
	Port int = 8080
)
