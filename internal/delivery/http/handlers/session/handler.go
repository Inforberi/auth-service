package session

import (
	"log/slog"

	"github.com/inforberi/auth-service/internal/service/auth/session"
)

type SessionHandler struct {
	sessionService *session.SessionService
	log            *slog.Logger
}

func New(sessionService *session.SessionService, log *slog.Logger) *SessionHandler {
	return &SessionHandler{sessionService: sessionService, log: log}
}
