package auth

import (
	"encoding/json"
	"net/http"

	helpers "github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/service/auth"
)

func (h *AuthHandler) RegisterEmail(w http.ResponseWriter, r *http.Request) {
	var req registerEmailRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid json body")
		return
	}

	// get client info
	client := helpers.ExtractClientInfo(r)

	input := auth.RegisterInput{
		Email:     req.Email,
		Password:  req.Password,
		IP:        client.IP,
		UserAgent: client.UserAgent,
		DeviceID:  client.DeviceID,
	}

	// service register
	res, err := h.authService.RegisterEmail(r.Context(), input)
	if err != nil {
		if status, code, message, ok := mapAuthError(err); ok {
			helpers.WriteError(w, status, code, message)
			return
		}

		// unknown errors
		h.log.Error("register email failed",
			"err", err,
			"email", req.Email,
			"ip", client.IP,
			"method", r.Method,
			"path", r.URL.Path,
		)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.SetSessionCookie(w, res.Token, res.ExpiresAt)

	helpers.WriteJSON(w, http.StatusCreated, registerEmailResponse{
		UserID: res.UserID,
	})
}
