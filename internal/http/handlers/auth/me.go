package auth

import (
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
)

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.ReadSessionCookie(r)
	if err != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing session")
		return
	}

	authInfo, err := h.authService.Me(r.Context(), token)
	if err != nil {
		if status, code, message, ok := mapAuthError(err); ok {
			helpers.WriteError(w, status, code, message)
			return
		}

		h.log.Error("me failed",
			"err", err,
			"path", r.URL.Path,
			"method", r.Method,
		)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, MeResponse{UserID: authInfo.UserID})

}
