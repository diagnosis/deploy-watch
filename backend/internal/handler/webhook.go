package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diagnosis/deploy-watch/internal/apperror"
	"github.com/diagnosis/deploy-watch/internal/helper"
	"github.com/diagnosis/deploy-watch/internal/logger"
	"github.com/diagnosis/deploy-watch/internal/sse"
	"github.com/diagnosis/deploy-watch/internal/store"
)

type GitHubPushEvent struct {
	Ref        string `json:"ref"` // "refs/heads/master"
	Repository struct {
		Name     string `json:"name"`      // "codecombat"
		FullName string `json:"full_name"` // "walkingtospace/codecombat"
	} `json:"repository"`
	Sender struct {
		ID    int64  `json:"id"`    // 2024264 ← GitHub user ID!
		Login string `json:"login"` // "walkingtospace"
	} `json:"sender"`
	HeadCommit struct {
		ID      string `json:"id"`      // "40a717b3644e..."
		Message string `json:"message"` // "Merge pull request #10..."
		Author  struct {
			Name string `json:"name"` // "Honam"
		} `json:"author"`
	} `json:"head_commit"`
}

type WebhookHandler struct {
	userStore   store.UserStore
	deployStore store.DeployStore
	broadcaster *sse.Broadcaster
}

func NewWebhookHandler(userStore store.UserStore, deployStore store.DeployStore, broadcaster *sse.Broadcaster) *WebhookHandler {
	return &WebhookHandler{
		userStore:   userStore,
		deployStore: deployStore,
		broadcaster: broadcaster,
	}
}
func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		logger.Error(ctx, "method not allowed", "method", r.Method)
		helper.RespondError(w, r, apperror.BadRequest("only POST allowed"))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var event GitHubPushEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		logger.Error(ctx, "failed to decode webhook", "err", err)
		helper.RespondError(w, r, apperror.BadRequest("invalid JSON"))
		return
	}

	// Parse branch
	branch := event.Ref
	if len(branch) > 11 && branch[:11] == "refs/heads/" {
		branch = branch[11:]
	}

	// Find user
	user, err := h.userStore.GetByGitHubID(ctx, event.Sender.ID)
	if err != nil {
		logger.Warn(ctx, "user not found for webhook", "github_id", event.Sender.ID)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Save to DB
	deployEvent, err := h.deployStore.Create(
		ctx,
		user.ID,
		event.Repository.Name,
		event.HeadCommit.ID,
		event.HeadCommit.Message,
		event.HeadCommit.Author.Name,
		branch,
		"success",
	)
	if err != nil {
		logger.Error(ctx, "failed to save deploy event", "err", err)
		helper.RespondError(w, r, apperror.InternalError("failed to save deploy", err))
		return
	}

	// Broadcast to user's SSE
	message := fmt.Sprintf("🚀 %s pushed to %s/%s - %s (%s)",
		deployEvent.Author,
		deployEvent.RepoName,
		deployEvent.Branch,
		deployEvent.CommitMessage,
		deployEvent.CommitSHA[:7],
	)

	h.broadcaster.BroadcastToUser(user.ID, message)

	logger.Info(ctx, "webhook processed", "repo", deployEvent.RepoName, "user_id", user.ID)
	helper.RespondMessage(w, r, 200, "ok")
}
