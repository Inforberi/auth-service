package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/service/auth"
)

func (h *AuthHandler) LoginEmail(w http.ResponseWriter, r *http.Request) {
	var req LoginEmailRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "invalid_login", "invalid json body")
		return
	}

	// get client info
	client := helpers.ExtractClientInfo(r)

	input := auth.LoginInput{
		Email:     req.Email,
		Password:  req.Password,
		IP:        client.IP,
		UserAgent: client.UserAgent,
		DeviceID:  client.DeviceID,
	}

	res, err := h.authService.LoginWithEmail(r.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrInvalidCredentials):
			helpers.WriteError(w, http.StatusUnauthorized, "invalid_credentials", err.Error())
			return

		case errors.Is(err, auth.ErrUserDisabled):
			helpers.WriteError(w, http.StatusForbidden, "user_disabled", err.Error())
			return

		default:
			h.log.Error("login email failed",
				"err", err,
				"email", req.Email,
				"method", r.Method,
				"path", r.URL.Path,
			)

			helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}
	}

	helpers.SetSessionCookie(w, res.Token, res.ExpiresAt)

	helpers.WriteJSON(w, http.StatusOK, LoginEmailResponse{
		UserID:    res.UserID,
		SessionID: res.SessionID,
		ExpiresAt: res.ExpiresAt.UTC().Format(time.RFC3339),
	})
}
