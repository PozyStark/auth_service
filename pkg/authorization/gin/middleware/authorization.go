package middleware

import (
	"auth-service/pkg/authorization"
	token "auth-service/pkg/authorization/gin/token"
	"auth-service/pkg/token/handler/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(
	auth *authorization.Authorization,
	tokenHandler *jwt.JWTTokenHandler,
) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token, err := token.GetAccessToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg":   "Something went wrong",
					"error": err.Error(),
				},
			)
			return
		}

		payload, err := tokenHandler.ParseAccessToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg":   "Something went wrong",
					"error": err.Error(),
				},
			)
			return
		}

		haveAccess := auth.HaveAccess(
			&authorization.UserAccess{
				Permissions: payload.Permissions,
				Roles:       payload.Roles,
				Groups:      payload.Groups,
			},
		)

		if !haveAccess {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"msg": "The user does not have access to the resource",
				},
			)
			return
		}
		ctx.Next()
	}
}
