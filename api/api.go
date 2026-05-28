package api

import (
	"auth-service/internal/config"
	db "auth-service/internal/db/gorm"
	"auth-service/internal/services"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	HEADER = "header"
	BODY   = "body"
	COOKIE = "cookie"
)

var (
	Config                *config.Config
	DbConnection          *db.GormDbConnection
	UserService           *services.UserService
	RegistryTokenService  *services.RegistryTokenService
	UserPermissionService *services.UserPermissionService
	UserRoleService       *services.UserRoleService
	UserGroupService      *services.UserGroupService
	Logger                *slog.Logger
	MainContext           *context.Context
)

func LoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestStart := time.Now()
		ctx.Next()
		status := ctx.Writer.Status()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		requestTime := time.Since(requestStart)
		logger.Info("Status: %s | Method: %s | Path: %s | ExecTime: %v\n", strconv.Itoa(status), method, path, requestTime)
		requestData := fmt.Sprintf("Headers: %v | Body: %v", ctx.Request.Header, ctx.Request.Body)
		logger.Debug("RequestData: " + requestData)
	}
}

func NewApiRouter(
	config *config.Config,
	userService *services.UserService,
	registryTokenService *services.RegistryTokenService,
	userPermissionService *services.UserPermissionService,
	userRoleService *services.UserRoleService,
	userGroupService *services.UserGroupService,
	logger *slog.Logger,
	mainContext *context.Context,
) (router *gin.Engine) {
	Config = config
	UserService = userService
	RegistryTokenService = registryTokenService
	UserPermissionService = userPermissionService
	UserRoleService = userRoleService
	UserGroupService = userGroupService
	Logger = logger
	MainContext = mainContext

	router = gin.Default()

	return router
}
