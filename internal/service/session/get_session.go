package session

import (
	"context"
)

type GetSessionResult struct {
	SessionID      string
	UserID         string
	SessionVersion int
}

func (s *SessionService) GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (GetSessionResult, error) {
	sessionID, userID, sessionVersion, expiresAt, revokedAt, actualSessionVersion, disabledAt, found, err := s.repo.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		return GetSessionResult{}, ErrGetSession
	}

	if !found {
		return GetSessionResult{}, ErrSessionNotFound
	}

	now := s.clock.Now().UTC()

	if revokedAt != nil && !revokedAt.After(now) {
		return GetSessionResult{}, ErrSessionIsRevoked
	}

	if !expiresAt.After(now) {
		return GetSessionResult{}, ErrSessionIsExpired
	}

	if actualSessionVersion != sessionVersion {
		return GetSessionResult{}, ErrSessionVersionMismatch
	}

	if disabledAt != nil && !disabledAt.After(now) {
		return GetSessionResult{}, ErrUserIsDisabled
	}

	return GetSessionResult{
		SessionID:      sessionID,
		UserID:         userID,
		SessionVersion: sessionVersion,
	}, nil
}
