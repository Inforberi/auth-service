package postgres

import (
	"context"
	"time"
)

func (s *AuthStore) CreateSession(ctx context.Context, userID string, sessionVersion int, tokenHash []byte, now, expiresAt time.Time, ip, ua *string) (sessionID string, err error) {

	err = s.db.QueryRow(ctx, `
		insert into sessions (
			user_id,
			session_version,
			token_hash,
			created_at,
			last_seen_at,
			expires_at,
			ip,
			user_agent
		)
		values ($1,$2,$3,$4,$4,$5,$6,$7)
		returning id
	`, userID, sessionVersion, tokenHash, now, expiresAt, ip, ua).Scan(&sessionID)

	if err != nil {
		return "", err
	}

	return sessionID, nil
}
