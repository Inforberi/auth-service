package auth

import (
	"log/slog"

	"github.com/inforberi/auth-service/internal/service/auth"
)

type AuthHandler struct {
	authService *auth.AuthService
	log         *slog.Logger
}

func NewAuthHandler(authService *auth.AuthService, log *slog.Logger) *AuthHandler {
	return &AuthHandler{authService: authService, log: log}
}
