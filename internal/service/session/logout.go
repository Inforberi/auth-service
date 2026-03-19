package session

import (
	"context"
	"fmt"
)

func (s *SessionService) Logout(ctx context.Context, sessionID string, tokenHash []byte) error {

	if err := s.revokeSession(ctx, sessionID); err != nil {
		return fmt.Errorf("%w: %v", ErrRevokeSession, err)
	}

	if s.cache != nil {
		if err := s.cache.MarkSessionRevoked(ctx, tokenHash, s.sessionTTL); err != nil {
			return fmt.Errorf("%w: %v", ErrCacheSync, err)
		}

		_ = s.cache.DeleteSession(ctx, tokenHash)
	}

	return nil
}

func (s *SessionService) LogoutAll(ctx context.Context, userID string) error {

	newVersion, err := s.repo.IncrementUserSessionVersion(ctx, userID, s.clock.Now().UTC())
	if err != nil {
		if isRepoUserNotFound(err) {
			return ErrUserNotFound
		}
		return fmt.Errorf("%w: %v", ErrLogoutAll, err)
	}

	if s.cache != nil {
		if err := s.cache.SetUserSessionVersion(ctx, userID, newVersion); err != nil {
			return fmt.Errorf("%w: %v", ErrCacheSync, err)
		}
	}
	
	return nil
}
