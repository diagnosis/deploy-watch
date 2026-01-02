package handler

import (
	"net/http"
	"os"

	"github.com/diagnosis/deploy-watch/internal/apperror"
	"github.com/diagnosis/deploy-watch/internal/auth"
	"github.com/diagnosis/deploy-watch/internal/helper"
	"github.com/diagnosis/deploy-watch/internal/logger"
	"github.com/diagnosis/deploy-watch/internal/middleware"
	"github.com/diagnosis/deploy-watch/internal/store"
)

type AuthHandler struct {
	oauth       *auth.GitHubOAuth
	userStore   store.UserStore
	deployStore store.DeployStore
}

func NewAuthHandler(oauth *auth.GitHubOAuth, userStore store.UserStore, deployStore store.DeployStore) *AuthHandler {
	return &AuthHandler{oauth: oauth, userStore: userStore, deployStore: deployStore}
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := auth.GenerateStateToken()
	if err != nil {
		logger.Error(r.Context(), "failed to generate state", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to generate state", err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth-state",
		Value:    state,
		Path:     "/",
		MaxAge:   600,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	authURL := h.oauth.GetAuthURL(state)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	stateCookie, err := r.Cookie("oauth-state")
	if err != nil {
		logger.Error(ctx, "oauth state cookie not found", "err", err)
		helper.RespondError(w, r, apperror.InternalError("internal server error", err))
		return
	}

	stateFromURL := r.URL.Query().Get("state")
	if stateFromURL == "" || stateCookie.Value != stateFromURL {
		if stateFromURL == "" {
			logger.Warn(ctx, "oauth state missing from callback URL")
		} else {
			logger.Warn(ctx, "oauth state mismatch")
		}
		helper.RespondError(w, r, apperror.BadRequest("invalid OAuth state"))
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		logger.Warn(ctx, "authorization code missing from callback")
		helper.RespondError(w, r, apperror.BadRequest("authorization code is missing"))
		return
	}

	token, err := h.oauth.ExchangeCode(ctx, code)
	if err != nil {
		logger.Error(ctx, "failed to exchange authorization code for token", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to complete OAuth flow", err))
		return
	}

	githubUser, err := h.oauth.GetGitHubUser(ctx, token)
	if err != nil {
		logger.Error(ctx, "failed to fetch GitHub user profile", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to retrieve user profile", err))
		return
	}

	regUser, err := h.userStore.GetByGitHubID(ctx, githubUser.ID)
	if err != nil {
		logger.Info(ctx, "user not found, creating new account", "github_id", githubUser.ID)

		regUser, err = h.userStore.Create(
			ctx,
			githubUser.ID,
			githubUser.Login,
			githubUser.Email,
			githubUser.AvatarURL,
			token.AccessToken,
		)
		if err != nil {
			logger.Error(ctx, "failed to create user account", "err", err)
			helper.RespondError(w, r, apperror.InternalError("failed to create user account", err))
			return
		}
	}

	auth.SetSessionCookie(w, regUser.ID)

	// Clear oauth state cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth-state",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // Default for dev
	}
	http.Redirect(w, r, frontendURL, http.StatusSeeOther)
}

func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	auth.ClearSessionCookie(w)
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // Default for dev
	}
	http.Redirect(w, r, frontendURL, http.StatusSeeOther)
}

func (h *AuthHandler) HandleMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c, err := auth.GetSessionCookie(r)
	if err != nil {
		logger.Error(ctx, "failed to get session cookie", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to get session cookie", err))
		return
	}
	user, err := h.userStore.GetByID(ctx, c)
	if err != nil {
		logger.Error(ctx, "failed to get user", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to get user", err))
		return
	}
	helper.RespondJSON(w, r, 200, map[string]any{
		"user": user,
	})
}

func (h *AuthHandler) HandleGetDeploys(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, ok := middleware.GetUser(ctx)
	if !ok {
		helper.RespondError(w, r, apperror.Unauthorized("not authorized"))
		return
	}
	deploys, err := h.deployStore.GetByUserID(ctx, user.ID, 50)
	if err != nil {
		logger.Error(ctx, "failed to get deploys", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to get deploys", err))
		return
	}

	helper.RespondJSON(w, r, 200, map[string]any{
		"deploys": deploys,
	})

}
