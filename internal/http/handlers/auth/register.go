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

		if errors.Is(err, auth.ErrEmailTaken) {
			helpers.WriteError(w, http.StatusConflict, "email_taken", err.Error())
			return
		}

		// все validation ошибки
		if errors.Is(err, auth.ErrEmptyEmail) ||
			errors.Is(err, auth.ErrInvalidEmail) ||
			errors.Is(err, auth.ErrPasswordTooShort) ||
			errors.Is(err, auth.ErrPasswordTooLong) ||
			errors.Is(err, auth.ErrPasswordNoLetter) ||
			errors.Is(err, auth.ErrPasswordNoDigit) {

			helpers.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
			return
		}

		// неизвестная ошибка
		h.log.Error("register email failed", "err", err)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.SetSessionCookie(w, res.Token, res.ExpiresAt)

	helpers.WriteJSON(w, http.StatusCreated, registerEmailResponse{
		UserID:    res.UserID,
		SessionID: res.SessionID,
		ExpiresAt: res.ExpiresAt.UTC().Format(time.RFC3339),
	})
}
