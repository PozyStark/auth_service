package api

import (
	"auth-service/internal/app"
	"auth-service/internal/models"
	token "auth-service/pkg/authorization/gin/token"
	"auth-service/pkg/token/handler/jwt"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RefreshHandler(
	accessTokenHandler *jwt.JWTTokenHandler,
	refreshTokenHandler *jwt.JWTTokenHandler,
) gin.HandlerFunc {

	op := "RefreshHandler"

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
				http.StatusUnauthorized,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		user, err := UserService.GetUserByUserId(ctx, tokenPayload.UserId)
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
		if !user.IsActive {
			Logger.Info(
				fmt.Sprintf(
					"Operation: %s UserId: %s IsActive: %v Msg: %s",
					op, user.ID, user.IsActive, "User is not active",
				),
			)
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "User is not active",
				},
			)
			return
		}

		registryToken, err := RegistryTokenService.GetByTokenId(ctx, tokenPayload.TokenId)
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
		if !registryToken.Active {
			Logger.Info(fmt.Sprintf("Operation: %s Msg: %s", op, "Token is not active"))
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "Token is not active",
				},
			)
			return
		}
		if registryToken.Jti != tokenPayload.Jti {
			Logger.Info(
				fmt.Sprintf(
					"Operation: %s PayloadJti: %s RegistryJti: %s Msg: %s", 
					op, tokenPayload.Jti, registryToken.Jti,  "Token JTI do not match",
				),
			)
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "Token JTI do not match",
				},
			)
			return
		}

		userInfo, errs := app.GetUserInfo(
			ctx,
			*UserPermissionService,
			*UserRoleService,
			*UserGroupService,
			tokenPayload.UserId,
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

		updateRegistryToken := models.RegistryToken{
			ID:        tokenPayload.TokenId,
			Jti:       uuid.NewString(),
			UpdatedAt: time.Now().UTC(),
			ExpireAt:  time.Now().UTC().Add(Config.Token.RefreshTokenLifetime),
		}

		accessToken, err := accessTokenHandler.CreateToken(
			jwt.NewAccessTokenPayload(
				tokenPayload.Issuer,
				tokenPayload.TokenId,
				updateRegistryToken.Jti,
				Config.Token.AccessTokenLifetime,
				tokenPayload.UserId,
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
				tokenPayload.Issuer,
				tokenPayload.TokenId,
				updateRegistryToken.Jti,
				Config.Token.RefreshTokenLifetime,
				tokenPayload.UserId,
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

		err = RegistryTokenService.UpdateByModel(
			&updateRegistryToken,
		)
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
			int(Config.Token.RefreshTokenLifetime),
			"/", "localhost", false, true,
		)

		ctx.Header("Authorization", "Bearer "+accessToken)

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		)

	}

}
