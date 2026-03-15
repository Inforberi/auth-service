package session

import (
	"context"
	"errors"
	"fmt"
)

func (s *SessionService) Logout(ctx context.Context, sessionID string) error {

	if err := s.RevokeSession(ctx, sessionID); err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return nil
		}
		return fmt.Errorf("%w: %v", ErrRevokeSession, err)
	}

	return nil
}

func (s *SessionService) LogoutAll(ctx context.Context, userID string) error {

	if err := s.repo.IncrementUserSessionVersion(ctx, userID, s.clock.Now().UTC()); err != nil {
		if isRepoUserNotFound(err) {
			return ErrUserNotFound
		}
		return fmt.Errorf("%w: %v", ErrLogoutAll, err)
	}
	return nil
}
