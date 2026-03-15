package session

import (
	"context"
	"time"
)

func (s *SessionRepo) CreateSession(ctx context.Context, userID string, sessionVersion int, tokenHash []byte, now, expiresAt time.Time, ip, ua *string, deviceID *string) (sessionID string, err error) {
	err = s.db.QueryRow(ctx, `
	insert into sessions (
			user_id,
			session_version,
			token_hash,
			created_at,
			last_seen_at,
			expires_at,
			ip,
			user_agent,
			device_id
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		returning id
`, userID,
		sessionVersion,
		tokenHash,
		now,
		now,
		expiresAt,
		ip,
		ua,
		deviceID,
	).Scan(&sessionID)

	if err != nil {
		return "", err
	}

	return sessionID, nil
}
