package email

import (
	"log/slog"

	"github.com/inforberi/auth-service/internal/service/auth/email"
)

type EmailHandler struct {
	authService *email.EmailService
	log         *slog.Logger
}

func New(authService *email.EmailService, log *slog.Logger) *EmailHandler {
	return &EmailHandler{authService: authService, log: log}
}
