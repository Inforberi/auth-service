package session

import (
	"context"
	"time"
)

func (s *SessionRepo) UpdateSessionActivity(
	ctx context.Context,
	sessionID string,
	now time.Time,
	expiresAt time.Time,
	threshold time.Time,
	refreshBefore time.Time,
) (bool, error) {
	tag, err := s.db.Exec(ctx, `
		update sessions
		set
			last_seen_at = $2,
			expires_at = $3
		where id = $1
		  and last_seen_at < $4
		  and expires_at <= $5
	`, sessionID, now, expiresAt, threshold, refreshBefore)
	if err != nil {
		return false, err
	}

	return tag.RowsAffected() > 0, nil
}
