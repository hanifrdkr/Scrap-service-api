package app_rest

import (
	"context"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"helicopter-hr/config"
	"helicopter-hr/database"
	"helicopter-hr/internal/app_rest/handler/handler_auth"
	"helicopter-hr/internal/app_rest/handler/handler_candidate"
	"helicopter-hr/internal/app_rest/handler/handler_ping"
	"helicopter-hr/internal/app_rest/middleware"
	"helicopter-hr/internal/app_rest/middleware/jwtx"
	"helicopter-hr/internal/app_rest/repositories/repo_auth"
	"helicopter-hr/internal/app_rest/repositories/repo_candidate"
	"helicopter-hr/internal/app_rest/repositories/repo_user"
	"helicopter-hr/internal/app_rest/service/service_auth"
	"helicopter-hr/internal/app_rest/service/service_candidate"
	"helicopter-hr/pkg/zapx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func Run(configPath string) {
	loggerZap := zapx.CreateLogger()
	zap.ReplaceGlobals(loggerZap)
	defer func(loggerZap *zap.Logger) {
		err := loggerZap.Sync()
		if err != nil {
			panic(err)
		}
	}(loggerZap)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic("Error loading config file")
	}

	mysqlDB, err := database.InitSqlConnection(cfg)
	if err != nil {
		loggerZap.Panic("err init database", zap.Error(err))
	}

	router := gin.Default()
	router.Use(middleware.CustomGinzap(loggerZap))
	router.Use(ginzap.CustomRecoveryWithZap(loggerZap, true, middleware.ErrorHandler))
	router.Use(middleware.RequestIDMiddleware(time.Duration(cfg.App.Timeout) * time.Second))

	handler_ping.NewHandlerPing(router)

	repoAuth := repo_auth.NewAuthRepository(mysqlDB)
	repoUser := repo_user.NewUserRepository(mysqlDB)
	repoCandidate := repo_candidate.NewCandidateRepository(mysqlDB)

	jwtAuthentication := jwtx.NewJWTAuthentication(cfg, repoAuth)

	serviceAuth := service_auth.NewAuthService(repoAuth, repoUser, jwtAuthentication)
	handler_auth.NewHandlerAuth(cfg, router, jwtAuthentication, serviceAuth)

	serviceCandidate := service_candidate.NewCandidateService(cfg, repoCandidate)
	handler_candidate.NewHandlerUser(cfg, router, jwtAuthentication, serviceCandidate)

	// Create a server with desired configurations
	server := &http.Server{
		Addr:    "0.0.0.0:" + cfg.App.Port,
		Handler: router,
	}

	// Now, set up the signal handling to catch SIGINT (Ctrl+C) and SIGTERM (kill)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		zap.L().Info(fmt.Sprintf("Server running at: %s", cfg.App.Port))
		if err := server.ListenAndServe(); err != nil {
			zap.L().Fatal("Server error", zap.Error(err))
		}
	}()

	// Block until a signal is received
	<-done

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		// close
		cancel()
	}()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Error during server shutdown:", zap.Error(err))
	}
	zap.L().Info("Server gracefully shut down")
}
