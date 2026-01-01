package middleware

import (
	"context"
	"net/http"

	"github.com/diagnosis/deploy-watch/internal/apperror"
	"github.com/diagnosis/deploy-watch/internal/auth"
	"github.com/diagnosis/deploy-watch/internal/helper"
	"github.com/diagnosis/deploy-watch/internal/logger"
	"github.com/diagnosis/deploy-watch/internal/store"
)

type ctxKey string

const UserCtxKey ctxKey = "user"

func RequireAuth(userStore store.UserStore) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := auth.GetSessionCookie(r)
			if err != nil {
				logger.Error(r.Context(), "unable to get session cookie", "err", err)
				helper.RespondError(w, r, apperror.Unauthorized("unauthorized"))
				return
			}
			user, err := userStore.GetByID(r.Context(), c)
			if err != nil {
				logger.Error(r.Context(), "unable to get user", "err", err)
				helper.RespondError(w, r, apperror.Unauthorized("unauthorized"))
				return
			}
			ctx := context.WithValue(r.Context(), UserCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func GetUser(ctx context.Context) (*store.User, bool) {
	user, ok := ctx.Value(UserCtxKey).(*store.User)
	return user, ok
}
