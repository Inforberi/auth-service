package session

import (
	"context"
	"errors"

	"github.com/inforberi/auth-service/internal/service/session"
	"github.com/jackc/pgx/v5"
)

func (s *SessionRepo) GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (
	session.GetRepoResult, error) {

	var res session.GetRepoResult

	err := s.db.QueryRow(ctx, `
		select
			s.id,
			s.user_id,
			s.session_version,
			s.expires_at,
			s.revoked_at,
			u.session_version,
			u.disabled_at
		from sessions s
		join users u on u.id = s.user_id
		where s.token_hash = $1
	`, tokenHash).Scan(
		&res.SessionID,
		&res.UserID,
		&res.SessionVersion,
		&res.ExpiresAt,
		&res.RevokedAt,
		&res.ActualSessionVersion,
		&res.DisabledAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return session.GetRepoResult{}, nil
		}
		return session.GetRepoResult{}, err
	}

	return session.GetRepoResult{SessionID: res.SessionID, UserID: res.UserID, SessionVersion: res.SessionVersion, ExpiresAt: res.ExpiresAt, RevokedAt: res.RevokedAt, ActualSessionVersion: res.ActualSessionVersion, DisabledAt: res.DisabledAt, Found: res.Found}, nil
}
