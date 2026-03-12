package session

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *SessionRepo) GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (
	sessionID string,
	userID string,
	sessionVersion int,
	expiresAt time.Time,
	revokedAt *time.Time,
	actualSessionVersion int,
	disabledAt *time.Time,
	found bool,
	err error,
) {
	err = s.db.QueryRow(ctx, `
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
		&sessionID,
		&userID,
		&sessionVersion,
		&expiresAt,
		&revokedAt,
		&actualSessionVersion,
		&disabledAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", 0, time.Time{}, nil, 0, nil, false, nil
		}
		return "", "", 0, time.Time{}, nil, 0, nil, false, err
	}

	return sessionID, userID, sessionVersion, expiresAt, revokedAt, actualSessionVersion, disabledAt, true, nil
}
