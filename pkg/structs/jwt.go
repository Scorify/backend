package structs

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type contextKey struct {
	name string
}

var USER_CTX_KEY = &contextKey{"username"}
