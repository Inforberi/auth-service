package session

import (
	"time"

	"github.com/inforberi/auth-service/internal/config"
)

type SessionService struct {
	repo                   SessionRepo
	clock                  Clock
	token                  TokenGenerator
	sessionTTL             time.Duration
	activityUpdateInterval time.Duration
}

func NewSessionService(repo SessionRepo, token TokenGenerator, clock Clock, auth *config.Auth) *SessionService {
	return &SessionService{repo: repo, token: token, clock: clock, sessionTTL: auth.SessionTTL, activityUpdateInterval: auth.UpdateInterval}
}
