package api

import (
	token "auth-service/pkg/authorization/gin/token"
	"auth-service/pkg/token/handler/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyHandler(
	accessTokenHandler *jwt.JWTTokenHandler,
) gin.HandlerFunc {

	op := "VerifyHandler"

	return func(ctx *gin.Context) {

		tokenString, err := token.GetAccessToken(ctx)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		_, err = accessTokenHandler.ParseAccessToken(tokenString)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"valid": false,
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"valid": true,
			},
		)

	}

}
