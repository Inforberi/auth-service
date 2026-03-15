package auth

import (
	"context"
	"crypto/sha256"
	"errors"

	"github.com/inforberi/auth-service/internal/service/session"
)

type AuthInfo struct {
	UserID         string
	SessionID      string
	SessionVersion int
}

func (s *AuthService) Me(ctx context.Context, token string) (AuthInfo, error) {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hash[:]

	sess, err := s.sessions.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		switch {
		case errors.Is(err, session.ErrSessionNotFound),
			errors.Is(err, session.ErrSessionIsRevoked),
			errors.Is(err, session.ErrSessionIsExpired),
			errors.Is(err, session.ErrSessionVersionMismatch),
			errors.Is(err, session.ErrUserIsDisabled):
			return AuthInfo{}, ErrUnauthorized
		default:
			return AuthInfo{}, err
		}
	}

	return AuthInfo{UserID: sess.UserID, SessionID: sess.SessionID, SessionVersion: sess.SessionVersion}, nil

}
