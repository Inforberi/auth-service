package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/inforberi/auth-service/internal/config"
	infraPG "github.com/inforberi/auth-service/internal/infra/postgres"
	repoPG "github.com/inforberi/auth-service/internal/repository/postgres"
	"github.com/inforberi/auth-service/internal/service/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Log    *slog.Logger
	pgPool *pgxpool.Pool
}

func NewApp(cfg *config.Config, log *slog.Logger) (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pgPool, err := infraPG.NewPgPool(ctx, cfg.Postgres)
	if err != nil {
		log.Error("failed to create pg pool", "err", err)
		return nil, err
	}

	repo := repoPG.NewAuthStore(pgPool)

	// service deps
	clock := auth.SystemClock{}
	hasher := auth.Argon2idHasher{}

	// TODO вытаскиваем зависимость, и кледм в handler
	_ = auth.NewService(repo, clock, hasher)

	return &App{
		Log:    log,
		pgPool: pgPool,
	}, nil
}
