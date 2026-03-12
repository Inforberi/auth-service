package session

import (
	"errors"
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
)

func mapSessionError(err error) (status int, code string, message string, ok bool) {
	if errors.Is(err, helpers.ErrSessionCookieNotFound) {
		return http.StatusUnauthorized, "unauthorized", "missing session", true
	}

	return 0, "", "", false
}
