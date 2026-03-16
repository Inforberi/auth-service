package session

import (
	"time"

	"github.com/inforberi/auth-service/internal/config"
)

type SessionService struct {
	repo                   SessionRepo
	cache                  SessionCache
	clock                  Clock
	token                  TokenGenerator
	sessionTTL             time.Duration
	activityUpdateInterval time.Duration
}

func NewSessionService(repo SessionRepo, token TokenGenerator, clock Clock, auth *config.Auth, cache SessionCache) *SessionService {
	return &SessionService{repo: repo, token: token, clock: clock, sessionTTL: auth.SessionTTL, activityUpdateInterval: auth.UpdateInterval, cache: cache}
}
