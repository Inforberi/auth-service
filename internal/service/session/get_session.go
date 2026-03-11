package session

import (
	"context"
)

func (s *SessionService) GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (GetSessionResult, error) {
	session, err := s.repo.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		return GetSessionResult{}, ErrGetSession
	}

	if !session.Found {
		return GetSessionResult{}, ErrSessionNotFound
	}

	now := s.clock.Now().UTC()

	if session.RevokedAt != nil && !session.RevokedAt.After(now) {
		return GetSessionResult{}, ErrSessionIsRevoked
	}

	if !session.ExpiresAt.After(now) {
		return GetSessionResult{}, ErrSessionIsExpired
	}

	if session.ActualSessionVersion != session.SessionVersion {
		return GetSessionResult{}, ErrSessionVersionMismatch
	}

	if session.DisabledAt != nil && !session.DisabledAt.After(now) {
		return GetSessionResult{}, ErrUserIsDisabled
	}

	return GetSessionResult{
		SessionID:      session.SessionID,
		UserID:         session.UserID,
		SessionVersion: session.SessionVersion,
	}, nil
}
