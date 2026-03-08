package session

import (
	"context"
	"errors"
)

func (s *SessionService) RevokeSession(ctx context.Context, sessionID string) error {
	now := s.clock.Now().UTC()
	err := s.repo.RevokeSession(ctx, sessionID, now)
	if err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return ErrSessionNotFound
		}
		return ErrRevokeSession
	}

	return nil
}
