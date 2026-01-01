package middleware

import (
	"net/http"

	"github.com/diagnosis/deploy-watch/internal/helper"
	"github.com/google/uuid"
)

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" {
			cid = uuid.NewString()
		}
		ctx := helper.WithCorrelationID(r.Context(), cid)
		w.Header().Set("X-Correlation-ID", cid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
