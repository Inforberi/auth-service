package auth

import (
	"encoding/json"
	"net/http"

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
		if status, code, message, ok := mapAuthError(err); ok {
			helpers.WriteError(w, status, code, message)
			return
		}

		h.log.Error("login email failed",
			"err", err,
			"email", req.Email,
			"method", r.Method,
			"path", r.URL.Path,
		)

		helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
		return
	}

	helpers.SetSessionCookie(w, res.Token, res.ExpiresAt)

	helpers.WriteJSON(w, http.StatusOK, LoginEmailResponse{
		UserID: res.UserID,
	})
}
