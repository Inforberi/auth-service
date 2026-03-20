package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type AuthRepo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{db: db}
}
