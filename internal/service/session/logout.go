package session

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
)

func (s *SessionService) Logout(ctx context.Context, token string) error {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hash[:]

	session, err := s.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return nil
		}
		return fmt.Errorf("%w: %v", ErrRevokeSession, err)
	}

	if err = s.RevokeSession(ctx, session.SessionID); err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return nil
		}
		return fmt.Errorf("%w: %v", ErrRevokeSession, err)
	}

	return nil
}

func (s *SessionService) LogoutAll(ctx context.Context, token string) error {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hash[:]

	session, err := s.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return nil
		}
		return fmt.Errorf("%w: %v", ErrLogoutAll, err)
	}

	if err = s.repo.IncrementUserSessionVersion(ctx, session.UserID, s.clock.Now().UTC()); err != nil {
		return fmt.Errorf("%w: %v", ErrLogoutAll, err)
	}
	return nil
}
