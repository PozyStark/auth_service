package pkg

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRefreshToken(ctx *gin.Context) (string, error) {
	cookieValue, err := ctx.Cookie("refresh_token")
	return cookieValue, err
}

func GetAccessToken(ctx *gin.Context) (string, error) {

	tokenStr := ctx.GetHeader("Authorization")

	if tokenStr == "" {
		return "", errors.New("Authorization header is required")
	}

	tokenParts := strings.SplitN(tokenStr, " ", 2)
	if !(len(tokenParts) == 2 && tokenParts[0] == "Bearer") {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return tokenParts[1], nil
}
