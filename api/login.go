package api

import (
	"auth-service/internal/app"
	"auth-service/internal/models"
	"auth-service/internal/schemas"
	"auth-service/pkg/cryptography"
	"auth-service/pkg/token/handler/jwt"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samborkent/uuidv7"
)

func LoginHandler(
	cryptParams *cryptography.CryptParams,
	accessTokenHandler *jwt.JWTTokenHandler,
	refreshTokenHandler *jwt.JWTTokenHandler,
) gin.HandlerFunc {

	op := "LoginHandler"

	return func(ctx *gin.Context) {

		var loginSchema schemas.Login

		if err := ctx.ShouldBind(&loginSchema); err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		user, err := UserService.GetUserByUsernameAndPassword(
			ctx,
			loginSchema.Username,
			cryptography.HashPassword(loginSchema.Password, *cryptParams),
		)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "Wrong username or password",
				},
			)
			return
		}

		userInfo, errs := app.GetUserInfo(
			ctx,
			*UserPermissionService,
			*UserRoleService,
			*UserGroupService,
			user.ID,
		)
		if len(errs) != 0 {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, errs))
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "Something went wrong",
				},
			)
			return
		}

		registryToken := models.RegistryToken{
			ID:        uuidv7.New().String(),
			UserID:    user.ID,
			Jti:       uuid.NewString(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			ExpireAt:  time.Now().UTC().Add(time.Duration(Config.Token.RefreshTokenLifetime)),
			Active:    true,
		}

		accessToken, err := accessTokenHandler.CreateToken(
			jwt.NewAccessTokenPayload(
				"localhost/auth-service",
				registryToken.ID,
				registryToken.Jti,
				time.Hour,
				user.ID,
				userInfo.Permissions,
				userInfo.Roles,
				userInfo.Groups,
			),
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

		refreshToken, err := refreshTokenHandler.CreateToken(
			jwt.NewRefreshTokenPayload(
				"localhost/auth-service",
				registryToken.ID,
				registryToken.Jti,
				time.Hour,
				user.ID,
			),
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

		err = RegistryTokenService.Save(&registryToken)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "Something went wrong",
				},
			)
			return
		}

		ctx.SetCookie(
			"refresh_token",
			refreshToken,
			int(Config.Token.RefreshTokenLifetime/time.Second),
			"/", "localhost", false, true,
		)
		ctx.Header("Authorization", "Bearer "+accessToken)

		ctx.JSON(http.StatusOK,
			gin.H{
				"refresh_token": refreshToken,
				"access_token":  accessToken,
			},
		)

	}
}
