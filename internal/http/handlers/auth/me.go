package auth

import (
	"errors"
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/service/auth"
)

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.ReadSessionCookie(r)
	if err != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing session")
		return
	}

	userID, err := h.authService.Me(r.Context(), token)
	if err != nil {
		if errors.Is(err, auth.ErrUnauthorized) {
			helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "invalid session")
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

	helpers.WriteJSON(w, http.StatusOK, MeResponse{UserID: userID})

}
