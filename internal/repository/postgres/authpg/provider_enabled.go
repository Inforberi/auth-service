package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *AuthRepo) IsProviderEnabled(ctx context.Context, code string) (bool, error) {
	var enabled bool

	err := s.db.QueryRow(ctx, `
		select enabled
		from auth_providers
		where code = $1
	`, code).Scan(&enabled)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// провайдера нет → считаем выключенным
			return false, nil
		}
		return false, err
	}

	return enabled, nil
}
