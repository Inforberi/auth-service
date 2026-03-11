package session

import (
	"context"
	"time"

	"github.com/inforberi/auth-service/internal/service/session"
)

func (s *SessionRepo) RevokeSession(
	ctx context.Context,
	sessionID string,
	now time.Time,
) error {

	ct, err := s.db.Exec(ctx, `
		update sessions
		set revoked_at = $2
		where id = $1
		  and revoked_at is null
	`, sessionID, now)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		var exists bool
		err = s.db.QueryRow(ctx, `
			select exists(select 1 from sessions where id = $1)
		`, sessionID).Scan(&exists)
		if err != nil {
			return err
		}
		if !exists {
			return session.ErrSessionNotFound
		}
		return nil
	}

	return nil
}

func (s *SessionRepo) IncrementUserSessionVersion(ctx context.Context, userID string, now time.Time) error {
	ct, err := s.db.Exec(ctx, `
	update users
	set session_version = session_version + 1
	updated_at = $2
	where id = $1
	`, userID, now)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return session.ErrUserNotFound
	}

	return nil
}
