package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/scorify/backend/pkg/config"
	"github.com/scorify/backend/pkg/data"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/ent/user"
	"github.com/scorify/backend/pkg/structs"
)

func JWTMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("auth")
	if err != nil {
		ctx.Next()
		return
	}

	claims := &structs.Claims{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWTKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.Next()
			return
		}
		ctx.Next()
		return
	}
	if !jwtToken.Valid {
		ctx.Next()
		return
	}

	entUser, err := data.Client.User.
		Query().
		Where(
			user.UsernameEQ(claims.Username),
		).
		Only(data.Ctx)
	if err != nil {
		ctx.Next()
		return
	}

	ctx.Request = ctx.Request.WithContext(
		context.WithValue(
			ctx.Request.Context(),
			structs.USER_CTX_KEY,
			entUser,
		),
	)

	ctx.Next()
}

func Parse(ctx context.Context) (*ent.User, error) {
	user, ok := ctx.Value(structs.USER_CTX_KEY).(*ent.User)
	if !ok {
		return nil, fmt.Errorf("invalid user")
	}
	return user, nil
}
