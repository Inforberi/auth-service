package session

import (
	"net/http"

	"github.com/inforberi/auth-service/internal/delivery/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/delivery/http/middleware"
)

func (s *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authInfo, ok := middleware.GetAuthContext(r.Context())
	if !ok {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing auth context")
		return
	}

	if err := s.sessionService.Logout(r.Context(), authInfo.SessionID, authInfo.TokenHash); err != nil {
		if status, code, message, ok := mapSessionError(err); ok {
			helpers.WriteError(w, status, code, message)
			return
		}

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
	authInfo, ok := middleware.GetAuthContext(r.Context())
	if !ok {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing auth context")
		return
	}

	if err := s.sessionService.LogoutAll(r.Context(), authInfo.UserID); err != nil {
		if status, code, message, ok := mapSessionError(err); ok {
			helpers.WriteError(w, status, code, message)
			return
		}

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
