package session

import (
	"context"
)

func (s *SessionService) revokeSession(ctx context.Context, sessionID string) error {
	now := s.clock.Now().UTC()
	err := s.repo.RevokeSession(ctx, sessionID, now)
	if err != nil {
		if isRepoSessionNotFound(err) {
			return ErrSessionNotFound
		}
		return ErrRevokeSession
	}

	return nil
}
