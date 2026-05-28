package api

import (
	"auth-service/internal/services"
	token "auth-service/pkg/authorization/gin/token"
	"auth-service/pkg/token/handler/jwt"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(
	tokenRegistryService *services.RegistryTokenService,
	refreshTokenHandler *jwt.JWTTokenHandler,
) gin.HandlerFunc {

	op := "LogoutHandler"

	return func(ctx *gin.Context) {

		tokenString, err := token.GetRefreshToken(ctx)
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

		tokenPayload, err := refreshTokenHandler.ParseRefreshToken(tokenString)
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

		err = tokenRegistryService.UpdateById(
			tokenPayload.TokenId,
			map[string]any{
				"active":     false,
				"updated_at": time.Now().UTC(),
			},
		)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"msg": "Something went wrong",
				},
			)
			return
		}

		ctx.SetCookie(
			"refresh_token",
			"", -1, "/", "localhost", false, true,
		)

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg": "logout success",
			},
		)
	}

}

func LogoutAllHandler(
	tokenRegistryService *services.RegistryTokenService,
	refreshTokenHandler *jwt.JWTTokenHandler,
) gin.HandlerFunc {

	op := "LogoutAllHandler"

	return func(ctx *gin.Context) {

		tokenString, err := token.GetRefreshToken(ctx)
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

		tokenPayload, err := refreshTokenHandler.ParseRefreshToken(tokenString)
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

		err = tokenRegistryService.UpdateByUserId(
			tokenPayload.UserId,
			map[string]any{
				"active":     false,
				"updated_at": time.Now().UTC(),
			},
		)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"msg": "Something went wrong",
				},
			)
			return
		}

		ctx.SetCookie(
			"refresh_token",
			"", -1, "/", "localhost", false, true,
		)

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg": "logout from all sessions success",
			},
		)

	}

}
