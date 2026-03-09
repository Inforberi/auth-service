package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/inforberi/auth-service/internal/service/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *AuthRepo) CreateUserWithEmailPassword(ctx context.Context, email, emailNorm, passwordHash string, now time.Time) (userID string, sessionVersion int, err error) {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", 0, err
	}
	// гарантируем rollback если не commit
	defer func() { _ = tx.Rollback(ctx) }()

	err = tx.QueryRow(ctx, `
		insert into users (created_at, updated_at)
		values ($1, $1)
		returning id, session_version
	`, now).Scan(&userID, &sessionVersion)
	if err != nil {
		return "", 0, err
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
			return "", 0, auth.ErrEmailTaken
		}
		return "", 0, err
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
		return "", 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", 0, err
	}

	return userID, sessionVersion, nil
}
