package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *AuthRepo) GetUserByEmail(
	ctx context.Context,
	emailNorm string,
) (userID string, passwordHash string, sessionVersion int, disabledAt *time.Time, found bool, err error) {

	var ph *string

	err = s.db.QueryRow(ctx, `
		select
		  u.id,
		  u.session_version,
		  u.disabled_at,
		  p.password_hash
		from user_identities i
		join users u on u.id = i.user_id
		left join user_passwords p on p.user_id = u.id
		where i.provider_code = 'email'
		  and i.identifier_normalized = $1
		limit 1
	`, emailNorm).Scan(&userID, &sessionVersion, &disabledAt, &ph)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", 0, nil, false, nil
		}
		return "", "", 0, nil, false, err
	}

	if ph != nil {
		passwordHash = *ph
	}

	return userID, passwordHash, sessionVersion, disabledAt, true, nil
}
