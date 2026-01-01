package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/diagnosis/deploy-watch/internal/application"
	"github.com/diagnosis/deploy-watch/internal/database"
	"github.com/diagnosis/deploy-watch/internal/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logger.Init()

	ctx := context.Background()
	dsn := os.Getenv("DB_URL_DEV")
	if os.Getenv("APP_ENV") == "prod" {
		dsn = os.Getenv("DB_URL_PROD")
	}

	pool, err := database.OpenPool(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	app := application.NewApplication(pool)
	router := app.SetupRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Info(ctx, "starting server 8080")
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(ctx, "server failed to start", "err", err)
			os.Exit(1)
		}
	}()
	logger.Info(ctx, "Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(ctx, "shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.Shutdown(shutdownCtx); err != nil {
		logger.Error(ctx, "server forced to shutdown", "err", err)
	}
	logger.Info(ctx, "server exit gracefully")

}
