package session

import (
	"time"

	"github.com/inforberi/auth-service/internal/config"
)

type SessionService struct {
	repo       SessionRepo
	cache      SessionCache
	clock      Clock
	token      TokenGenerator
	sessionTTL time.Duration
	auth       *config.Auth
}

func New(repo SessionRepo, token TokenGenerator, clock Clock, auth *config.Auth, cache SessionCache) *SessionService {
	return &SessionService{repo: repo, token: token, clock: clock, sessionTTL: auth.SessionTTL, auth: auth, cache: cache}
}
