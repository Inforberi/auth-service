package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/inforberi/auth-service/internal/config"
	authHandler "github.com/inforberi/auth-service/internal/http/handlers/auth"
	sessionHandler "github.com/inforberi/auth-service/internal/http/handlers/session"
	router "github.com/inforberi/auth-service/internal/http/router"
	"github.com/inforberi/auth-service/internal/infra/postgres"
	"github.com/inforberi/auth-service/internal/infra/redis"
	"github.com/inforberi/auth-service/internal/pkg"
	repoAuth "github.com/inforberi/auth-service/internal/repository/postgres/auth"
	repoSession "github.com/inforberi/auth-service/internal/repository/postgres/session"
	"github.com/inforberi/auth-service/internal/service/auth"
	"github.com/inforberi/auth-service/internal/service/session"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Log         *slog.Logger
	pgPool      *pgxpool.Pool
	redisClient *redis.Client
	Router      http.Handler
}

func NewApp(cfg *config.Config, log *slog.Logger) (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// postgres
	pgPool, err := postgres.NewPgPool(ctx, &cfg.Postgres)
	if err != nil {
		log.Error("failed to create pg pool", "err", err)
		return nil, err
	}

	// redis
	redis, err := redis.NewClient(ctx, cfg.Redis)
	if err != nil {
		log.Error("failed to create redis db", "err", err)
		return nil, err
	}

	// repo
	authRepo := repoAuth.NewAuthRepo(pgPool)
	sessionRepo := repoSession.NewSessionRepo(pgPool)

	// service deps
	clock := pkg.SystemClock{}
	hasher := auth.Argon2idHasher{}
	token := pkg.SecureTokenGenerator{}

	// services
	sessionService := session.NewSessionService(sessionRepo, token, clock, &cfg.Auth)
	authService := auth.NewAuthService(authRepo, clock, hasher, sessionService)

	// handlers
	authHandler := authHandler.NewAuthHandler(authService, log)
	sessionHandler := sessionHandler.NewSessionHandler(sessionService, log)

	// router
	router := router.NewRouter(authHandler, sessionHandler, authService, &cfg.HTTP)

	return &App{
		Log:         log,
		pgPool:      pgPool,
		Router:      router,
		redisClient: redis,
	}, nil
}
