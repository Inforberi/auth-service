package session

import (
	"context"
	"time"
)

func (s *SessionRepo) RevokeSession(
	ctx context.Context,
	sessionID string,
	now time.Time,
) error {

	// ВАЖНО: обновляем только если ещё не revoked (идемпотентность)
	ct, err := s.db.Exec(ctx, `
		update sessions
		set revoked_at = $2
		where id = $1
		  and revoked_at is null
	`, sessionID, now)
	if err != nil {
		return err
	}

	// Если 0 строк обновлено: либо сессии нет, либо она уже revoked.
	// Отличаем эти случаи, если тебе важно.
	if ct.RowsAffected() == 0 {
		// Проверим, существует ли сессия вообще.
		var exists bool
		err = s.db.QueryRow(ctx, `
			select exists(select 1 from sessions where id = $1)
		`, sessionID).Scan(&exists)
		if err != nil {
			return err
		}
		if !exists {
			return ErrSessionNotFound
		}
		// exists == true -> уже revoked, считаем успехом
		return nil
	}

	return nil
}
