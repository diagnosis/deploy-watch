package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/cors"
)

func CorsHandler() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			allowedOrigins := []string{
				"http://localhost:3000",
				"http://localhost:5173",
				"http://localhost:8080",
				"https://deploy-watch.vercel.app",
			}

			for _, allowed := range allowedOrigins {
				if origin == allowed {
					return true
				}
			}

			if strings.HasSuffix(origin, ".vercel.app") {
				return true
			}

			return false
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
