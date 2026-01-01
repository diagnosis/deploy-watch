package handler

import (
	"fmt"
	"net/http"

	"github.com/diagnosis/deploy-watch/internal/apperror"
	"github.com/diagnosis/deploy-watch/internal/helper"
	"github.com/diagnosis/deploy-watch/internal/logger"
	"github.com/diagnosis/deploy-watch/internal/middleware"
	"github.com/diagnosis/deploy-watch/internal/sse"
)

type SSEHandler struct {
	broadcaster *sse.Broadcaster
}

func NewSSEHandler(broadcaster *sse.Broadcaster) *SSEHandler {
	return &SSEHandler{broadcaster: broadcaster}
}

func (h *SSEHandler) HandleSSE(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := middleware.GetUser(ctx)
	if !ok {
		helper.RespondError(w, r, apperror.Unauthorized("not authorized"))
		logger.Error(ctx, "user not authorized")
		return
	}
	//headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")

	flusher, ok := w.(http.Flusher)
	if !ok {
		logger.Error(ctx, "stream not supported")
		helper.RespondError(w, r, apperror.BadRequest("stream not supported"))
		return
	}
	client := &sse.Client{
		UserID: user.ID,
		Send:   make(chan string),
	}

	h.broadcaster.Register(client)
	defer h.broadcaster.Unregister(client)

	for {
		select {
		case <-ctx.Done():
			logger.Warn(ctx, "channel closed")
			return
		case event := <-client.Send:
			fmt.Fprintf(w, "data: %s\n\n", event)
			flusher.Flush()
		}
	}

}
