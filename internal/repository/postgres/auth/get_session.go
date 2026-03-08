package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *AuthRepo) GetSessionByTokenHash(
	ctx context.Context,
	tokenHash []byte,
) (sessionID string, userID string, sessionVersion int, expiresAt time.Time, revokedAt *time.Time, err error) {

	err = s.db.QueryRow(ctx, `
		select id, user_id, session_version, expires_at, revoked_at
		from sessions
		where token_hash = $1
	`).Scan(&sessionID, &userID, &sessionVersion, &expiresAt, &revokedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", 0, time.Time{}, nil, nil
		}
		return "", "", 0, time.Time{}, nil, err
	}

	return sessionID, userID, sessionVersion, expiresAt, revokedAt, nil
}
