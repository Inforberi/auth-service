package auth

import (
	"log/slog"

	"github.com/inforberi/auth-service/internal/service/auth/email"
)

type AuthHandler struct {
	authService *email.EmailService
	log         *slog.Logger
}

func NewAuthHandler(authService *email.EmailService, log *slog.Logger) *AuthHandler {
	return &AuthHandler{authService: authService, log: log}
}
