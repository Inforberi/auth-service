package session

import (
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
)

func (s *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.ReadSessionCookie(r)
	if err != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing session")
		return
	}

	err = s.sessionService.RevokeSession(r.Context(), token)
	if err != nil {
		s.log.Error("logout failed",
			"err", err,
			"path", r.URL.Path,
			"method", r.Method,
		)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.ClearSessionCookie(w)
}
