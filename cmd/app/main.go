package main

import (
	"auth-service/api"
	"auth-service/internal/config"
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	repository "auth-service/internal/repository/gorm"
	"auth-service/internal/services"
	"auth-service/pkg/authorization"
	"auth-service/pkg/authorization/gin/middleware"
	"auth-service/pkg/cryptography"
	"auth-service/pkg/token/handler/jwt"
	"context"
	"log/slog"
	"os"
	"runtime"

	"gorm.io/gorm"
)

func main() {

	runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(1)

	mainContext := context.Background()
	cfg := config.LoadConfig()
	log := setupLogger(cfg.Env, cfg.Loglevel)
	// log = log.With(slog.String("env", cfg.Env))
	// log.Info("Initialized server",
	// 	slog.String("logLevel", cfg.Loglevel),
	// 	slog.String("address", cfg.HTTPServer.HttpAddress),
	// 	slog.String("port", strconv.Itoa(cfg.HTTPServer.HttpPort)),
	// 	slog.String("env", cfg.Env),
	// 	slog.String("DbHost", cfg.DbConfig.DbHost),
	// 	slog.String("DbName", cfg.DbConfig.DbName),
	// 	slog.String("DbUser", cfg.DbConfig.DbUser),
	// 	slog.String("SslMode", cfg.DbConfig.SslMode),
	// 	slog.String("EnableAutomigration", strconv.FormatBool(cfg.EnableAutomigration)),
	// )

	dbConnection, err := db.NewGormDbConnectionWithConfig(
		db.GetPostgressDialect(cfg.DbConfig),
		&gorm.Session{},
		&gorm.Config{},
	).Open()

	if err != nil {
		panic(err)
	}

	if cfg.EnableAutomigration {
		db.MustMakeMigrations(
			dbConnection,
			&gorm.Session{},
			models.User{},
			models.Role{},
			models.Group{},
			models.UserRole{},
			models.UserGroup{},
			models.UserPermission{},
			models.RolePermission{},
			models.GroupPermission{},
			models.RegistryToken{},
		)
	}

	userRepository := repository.NewGormUserRepository(dbConnection)
	registryTokenRepository := repository.NewGormRegistryTokenRepository(dbConnection)
	userPermissionRepository := repository.NewGormUserPermissionRepository(dbConnection)
	userRoleRepository := repository.NewGormUserRoleRepository(dbConnection)
	userGroupRepository := repository.NewGormUserGroupRepository(dbConnection)

	userService := services.NewUserService(userRepository)
	registyTokenService := services.NewRegistryTokenService(registryTokenRepository)
	userPermissionService := services.NewUserPermissionService(userPermissionRepository)
	userRoleService := services.NewUserRoleService(userRoleRepository)
	userGroupService := services.NewUserGroupService(userGroupRepository)

	jwtTokenHandlerAccess := jwt.NewTokenCreator(
		cfg.Token.SecretAccessToken, jwt.SIGN_METHOD_HS256,
	)
	jwtTokenHandlerRefresh := jwt.NewTokenCreator(
		cfg.Token.SecretRefreshToken,
		jwt.SIGN_METHOD_HS256,
	)

	router := api.NewApiRouter(
		cfg,
		userService,
		registyTokenService,
		userPermissionService,
		userRoleService,
		userGroupService,
		log,
		&mainContext,
	)

	cpyptParams := cryptography.CryptParams(cfg.Crypt)

	router.POST(
		"/registration",
		api.RegistrationHandler(&cpyptParams),
	)

	router.POST(
		"/login",
		api.LoginHandler(&cpyptParams, jwtTokenHandlerAccess, jwtTokenHandlerRefresh),
	)

	router.POST(
		"refresh",
		api.RefreshHandler(jwtTokenHandlerAccess, jwtTokenHandlerRefresh),
	)

	router.GET(
		"/verify",
		api.LoggerMiddleware(log),
		middleware.AuthMiddleware(&authorization.Authorization{}, jwtTokenHandlerAccess),
		api.VerifyHandler(jwtTokenHandlerAccess),
	)

	router.POST(
		"/logout",
		api.LogoutHandler(registyTokenService, jwtTokenHandlerRefresh),
	)

	router.POST(
		"/logout/all",
		api.LogoutAllHandler(registyTokenService, jwtTokenHandlerRefresh),
	)

	router.Run()

}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

const (
	logLevelInfo  = "info"
	logLevelWarn  = "warn"
	loglevelError = "error"
	logLevelDebug = "debug"
)

func setupLogger(env string, logLevel string) *slog.Logger {

	var log *slog.Logger
	var options slog.HandlerOptions

	switch logLevel {
	case logLevelInfo:
		options = slog.HandlerOptions{Level: slog.LevelInfo}
	case logLevelWarn:
		options = slog.HandlerOptions{Level: slog.LevelWarn}
	case loglevelError:
		options = slog.HandlerOptions{Level: slog.LevelError}
	case logLevelDebug:
		options = slog.HandlerOptions{Level: slog.LevelDebug}
	}

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &options))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &options))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &options))
	}

	return log
}
