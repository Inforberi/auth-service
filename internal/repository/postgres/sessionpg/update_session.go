package session

import (
	"context"
	"time"
)

func (s *SessionRepo) UpdateSessionActivity(ctx context.Context, sessionID string, now time.Time, expiresAt time.Time, threshold time.Time) error {

	_, err := s.db.Exec(ctx, `
		update sessions
		set 
			last_seen_at = $2,
			expires_at = $3
		where id = $1
		and last_seen_at < 4$
	`, sessionID, now, expiresAt, threshold)

	if err != nil {
		return err
	}

	return nil
}
