package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/inforberi/auth-service/internal/config"
	router "github.com/inforberi/auth-service/internal/http"
	authHandler "github.com/inforberi/auth-service/internal/http/handlers/auth"
	"github.com/inforberi/auth-service/internal/infra/postgres"
	"github.com/inforberi/auth-service/internal/pkg"
	repoAuth "github.com/inforberi/auth-service/internal/repository/postgres/auth"
	repoSession "github.com/inforberi/auth-service/internal/repository/postgres/session"
	"github.com/inforberi/auth-service/internal/service/auth"
	"github.com/inforberi/auth-service/internal/service/session"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Log    *slog.Logger
	pgPool *pgxpool.Pool
	Router http.Handler
}

func NewApp(cfg *config.Config, log *slog.Logger) (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pgPool, err := postgres.NewPgPool(ctx, cfg.Postgres)
	if err != nil {
		log.Error("failed to create pg pool", "err", err)
		return nil, err
	}

	authRepo := repoAuth.NewAuthRepo(pgPool)
	sessionRepo := repoSession.NewSessionRepo(pgPool)

	// service deps
	clock := pkg.SystemClock{}
	hasher := auth.Argon2idHasher{}
	token := pkg.SecureTokenGenerator{}

	// services
	sessionService := session.NewSessionService(sessionRepo, token, clock, cfg.Auth)
	authService := auth.NewAuthService(authRepo, clock, hasher, sessionService)

	// handler
	authHandler := authHandler.NewAuthHandler(authService)

	// router
	router := router.NewRouter(authHandler)

	return &App{
		Log:    log,
		pgPool: pgPool,
		Router: router,
	}, nil
}
