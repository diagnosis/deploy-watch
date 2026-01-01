package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/diagnosis/deploy-watch/internal/apperror"
	"github.com/diagnosis/deploy-watch/internal/helper"
)

var globalLogger *slog.Logger

func Init() {
	env := os.Getenv("APP_ENV")
	var handler slog.Handler
	if env == "production" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}
	globalLogger = slog.New(handler)
}
func Get() *slog.Logger {
	return globalLogger
}

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return helper.WithCorrelationID(ctx, correlationID)
}
func GetCorrelationId(ctx context.Context) string {
	return helper.GetCorrelationID(ctx)
}
func FromContext(ctx context.Context) *slog.Logger {
	logger := globalLogger
	if id := GetCorrelationId(ctx); id != "" {
		logger = logger.With("correlation_id", id)
	}
	return logger
}

func Info(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).InfoContext(ctx, msg, args...)
}
func Error(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).ErrorContext(ctx, msg, args...)
}
func Debug(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).DebugContext(ctx, msg, args...)
}
func Warn(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).WarnContext(ctx, msg, args...)
}

func LogError(ctx context.Context, err error) {
	ae := apperror.AsAppError(err)
	switch ae.Code {
	case apperror.CodeValidationError,
		apperror.CodeBadRequest:
		Warn(ctx, "client error", "code", ae.Code, "err", ae.Err)
	default:
		Error(ctx, "server error", "code", ae.Code, "err", ae.Err)
	}
}
