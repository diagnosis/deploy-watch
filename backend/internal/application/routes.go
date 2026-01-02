package application

import (
	"net/http"
	"time"

	"github.com/diagnosis/deploy-watch/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func (app *Application) SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.CorrelationMiddleware)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CorsHandler())

	// Public routes with timeout
	r.Group(func(r chi.Router) {
		r.Use(chimiddleware.Timeout(60 * time.Second))

		r.Get("/auth/github/login", app.authHandler.HandleLogin)
		r.Get("/auth/github/callback", app.authHandler.HandleCallback)
		r.Get("/auth/logout", app.authHandler.HandleLogout)
		r.Post("/webhook", app.webhookHandler.HandleWebhook)
	})

	// Protected routes with timeout (except SSE)
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth(app.userStore))
		r.Use(chimiddleware.Timeout(60 * time.Second))

		r.Get("/api/me", app.authHandler.HandleMe)
		r.Get("/api/deploys", app.authHandler.HandleGetDeploys)
	})

	// SSE endpoint without timeout
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth(app.userStore))

		r.Get("/api/events", app.sseHandler.HandleSSE)
	})

	r.Post("/test/broadcast", func(w http.ResponseWriter, r *http.Request) {
		app.broadcaster.Broadcast("🚀 Test broadcast message!")
		w.Write([]byte("Broadcasted!"))
	})

	return r

}
