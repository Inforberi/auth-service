package session

import (
	"errors"
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/service/auth/session"
)

func mapSessionError(err error) (status int, code string, message string, ok bool) {
	if errors.Is(err, helpers.ErrSessionCookieNotFound) {
		return http.StatusUnauthorized, "unauthorized", "missing session", true
	}
	if errors.Is(err, session.ErrUserNotFound) {
		return http.StatusUnauthorized, "unauthorized", "invalid session", true
	}

	return 0, "", "", false
}
