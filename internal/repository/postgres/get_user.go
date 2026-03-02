package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *AuthStore) GetUserIdByEmail(ctx context.Context, email string) (string, bool, error) {
	var userID string

	err := s.db.QueryRow(ctx, `
		select user_id
		from user_identities
		where provider_code = 'email' and identifier_normalized = $1
	`, email).Scan(&userID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", false, nil
		}
		return "", false, err
	}

	return userID, true, nil
}
