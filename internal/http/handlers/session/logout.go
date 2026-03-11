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

	if err = s.sessionService.Logout(r.Context(), token); err != nil {
		s.log.Error("logout failed",
			"err", err,
			"path", r.URL.Path,
			"method", r.Method,
		)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.ClearSessionCookie(w)
	w.WriteHeader(http.StatusNoContent)
}

func (s *SessionHandler) LogoutAll(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.ReadSessionCookie(r)
	if err != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing session")
		return
	}

	if err = s.sessionService.LogoutAll(r.Context(), token); err != nil {
		s.log.Error("logout all failed",
			"err", err,
			"path", r.URL.Path,
			"method", r.Method,
		)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.ClearSessionCookie(w)
	w.WriteHeader(http.StatusNoContent)
}
