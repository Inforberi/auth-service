package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *AuthStore) CreateUserWithEmailPassword(ctx context.Context, email, emailNorm, passwordHash string, now time.Time) (string, error) {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}
	// гарантируем rollback если не commit
	defer func() { _ = tx.Rollback(ctx) }()

	var userID string
	err = tx.QueryRow(ctx, `
		insert into users (created_at, updated_at)
		values ($1, $1)
		returning id
	`, now).Scan(&userID)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec(ctx, `
		insert into user_identities (
			user_id,
			provider_code,
			identifier,
			identifier_normalized,
			created_at,
			updated_at
		)
		values ($1, 'email', $2, $3, $4, $4)
	`, userID, email, emailNorm, now)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", ErrEmailTaken
		}
		return "", err
	}

	_, err = tx.Exec(ctx, `
		insert into user_passwords (
			user_id,
			password_hash,
			updated_at
		)
		values ($1, $2, $3)
	`, userID, passwordHash, now)

	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}

	return userID, nil
}
