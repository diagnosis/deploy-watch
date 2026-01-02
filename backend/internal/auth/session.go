package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func GenerateStateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func SetSessionCookie(w http.ResponseWriter, userID uuid.UUID) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_cookie",
		Value:    userID.String(),
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   os.Getenv("APP_ENV") == "prod",
	})
}
func GetSessionCookie(r *http.Request) (uuid.UUID, error) {
	cookie, err := r.Cookie("session_cookie")
	if err != nil {
		return uuid.Nil, err
	}
	id, err := uuid.Parse(cookie.Value)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_cookie",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
	})
}
