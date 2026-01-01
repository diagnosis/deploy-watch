package helper

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/diagnosis/deploy-watch/internal/apperror"
)

type ctxKey string

const correlationIDKey ctxKey = "correlation_id"

type ErrorResponse struct {
	Error struct {
		Code          string    `json:"code"`
		Message       string    `json:"message"`
		CorrelationID string    `json:",omitempty"`
		Timestamp     time.Time `json:"timestamp"`
	} `json:"error"`
}

type SuccessResponse struct {
	Data          any       `json:"data,omitempty"`
	Message       string    `json:"message,omitempty"`
	CorrelationID string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp"`
}

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIDKey, correlationID)
}
func GetCorrelationID(ctx context.Context) string {
	if id, ok := ctx.Value(correlationIDKey).(string); ok {
		return id
	}
	return ""
}

func RespondError(w http.ResponseWriter, r *http.Request, err error) {
	ctx := r.Context()
	correlationID := GetCorrelationID(ctx)
	//apperr instance
	ae := apperror.AsAppError(err)

	errorResponse := ErrorResponse{}
	errorResponse.Error.Code = string(ae.Code)
	errorResponse.Error.CorrelationID = correlationID
	errorResponse.Error.Message = ae.Message
	errorResponse.Error.Timestamp = time.Now().UTC()

	encodeJSON(w, ae.HTTPStatus, errorResponse)
}

func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data any) {
	ctx := r.Context()
	correlationID := GetCorrelationID(ctx)
	successResponse := SuccessResponse{
		Data:          data,
		CorrelationID: correlationID,
		Timestamp:     time.Now().UTC(),
	}
	encodeJSON(w, status, successResponse)
}

func RespondMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	ctx := r.Context()
	correlationID := GetCorrelationID(ctx)

	response := SuccessResponse{
		Message:       message,
		CorrelationID: correlationID,
		Timestamp:     time.Now().UTC(),
	}
	encodeJSON(w, status, response)
}

func encodeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
