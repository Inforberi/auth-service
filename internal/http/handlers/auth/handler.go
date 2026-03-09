package auth

import "github.com/inforberi/auth-service/internal/service/auth"

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
