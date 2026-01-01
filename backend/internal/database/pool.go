package database

import (
	"context"
	"time"

	"github.com/diagnosis/deploy-watch/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenPool(dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	cfg.MinConns = 2
	cfg.MaxConns = 10
	cfg.MaxConnLifetime = 25 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second
	cfg.ConnConfig.ConnectTimeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	logger.Info(ctx, "Connecting to db...")
	return pool, nil
}
