package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	helpers "github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/service/auth"
)

func (h *AuthHandler) RegisterEmail(w http.ResponseWriter, r *http.Request) {
	var req auth.RegisterInput

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid json body")
		return
	}

	client := helpers.ExtractClientInfo(r)

	input := auth.RegisterInput{
		Email:     req.Email,
		Password:  req.Password,
		IP:        client.IP,
		UserAgent: client.UserAgent,
		DeviceID:  client.DeviceID,
	}

	res, err := h.authService.RegisterEmail(r.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrProviderNotEnabled):
			helpers.WriteError(w, http.StatusBadRequest, "provider_not_enabled", err.Error())
		case errors.Is(err, auth.ErrEmailTaken):
			helpers.WriteError(w, http.StatusConflict, "email_taken", err.Error())
		default:
			helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		}
		return
	}

	helpers.SetSessionCookie(w, res.Token, res.ExpiresAt)

	helpers.WriteJSON(w, http.StatusCreated, registerEmailResponse{
		UserID:    res.UserID,
		SessionID: res.SessionID,
		ExpiresAt: res.ExpiresAt.UTC().Format(time.RFC3339),
	})
}
