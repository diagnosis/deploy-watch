package application

import (
	"github.com/diagnosis/deploy-watch/internal/auth"
	"github.com/diagnosis/deploy-watch/internal/handler"
	"github.com/diagnosis/deploy-watch/internal/sse"
	"github.com/diagnosis/deploy-watch/internal/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	pool           *pgxpool.Pool
	userStore      store.UserStore
	deployStore    store.DeployStore
	oauth          *auth.GitHubOAuth
	authHandler    *handler.AuthHandler
	broadcaster    *sse.Broadcaster
	sseHandler     *handler.SSEHandler
	webhookHandler *handler.WebhookHandler
}

func NewApplication(pool *pgxpool.Pool) *Application {
	//
	userStore := store.NewPGUserStore(pool)
	deployStore := store.NewPGDeployStore(pool)
	oauth := auth.NewGitHubOAuth()
	authHandler := handler.NewAuthHandler(oauth, userStore, deployStore)

	broadcaster := sse.NewBroadcaster()
	go broadcaster.Run()
	sseHandler := handler.NewSSEHandler(broadcaster)
	webhookHandler := handler.NewWebhookHandler(userStore, deployStore, broadcaster)
	return &Application{
		pool:           pool,
		userStore:      userStore,
		deployStore:    deployStore,
		oauth:          oauth,
		authHandler:    authHandler,
		broadcaster:    broadcaster,
		sseHandler:     sseHandler,
		webhookHandler: webhookHandler,
	}
}
