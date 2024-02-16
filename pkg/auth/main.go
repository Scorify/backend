package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/structs"
)

func GenerateJWT(username string, id uuid.UUID, role string) (string, int, error) {
	expiration := time.Now().Add(time.Duration(config.JWT.Timeout) * time.Hour)

	claims := &structs.Claims{
		Username: username,
		Role:     role,
		ID:       id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.JWT.Secret))

	return tokenStr, int(expiration.Unix()), err
}

func ParseJWT(tokenString string) (*jwt.Token, *structs.Claims, error) {
	claims := &structs.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})
	return token, claims, err
}
