package email

import (
	"context"
	"crypto/sha256"
)

type AuthInfo struct {
	UserID         string
	SessionID      string
	SessionVersion int
	TokenHash      []byte
}

func (s *EmailService) Me(ctx context.Context, token string) (AuthInfo, error) {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hash[:]

	sess, err := s.sessions.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		if isUnauthorizedSessionError(err) {
			return AuthInfo{}, ErrUnauthorized
		}
		return AuthInfo{}, err
	}

	_ = s.sessions.UpdateSessionActivity(ctx, sess.SessionID, tokenHash)

	return AuthInfo{UserID: sess.UserID, SessionID: sess.SessionID, SessionVersion: sess.SessionVersion, TokenHash: tokenHash}, nil

}
