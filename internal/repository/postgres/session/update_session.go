package session

import (
	"context"
	"time"
)

func (s *SessionRepo) UpdateSessionLastSeen(ctx context.Context, sessionID string, now time.Time) error {

	_, err := s.db.Exec(ctx, `
		update sessions
		where id = $1
		last_seen_at = $2
	`, sessionID, now)

	if err != nil {
		return err
	}

	return nil
}
