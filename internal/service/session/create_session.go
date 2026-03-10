package session

import (
	"context"
	"crypto/sha256"
)

func (s *SessionService) CreateSession(ctx context.Context, userID string, sessionVersion int, ip, ua, deviceID *string) (CreateSessionResult, error) {
	rawToken, err := s.token.New()
	if err != nil {
		return CreateSessionResult{}, ErrGenerateToken
	}

	hash := sha256.Sum256([]byte(rawToken))
	tokenHash := hash[:]

	now := s.clock.Now().UTC()
	expiresAt := now.Add(s.sessionTTL)

	sessionID, err := s.repo.CreateSession(ctx, userID, sessionVersion, tokenHash, now, expiresAt, ip, ua, deviceID)

	if err != nil {
		return CreateSessionResult{}, ErrCreateSession
	}

	return CreateSessionResult{
		SessionID: sessionID,
		Token:     rawToken,
		ExpiresAt: expiresAt,
	}, nil
}
