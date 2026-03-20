package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/delivery/http/router"

	emailHandler "github.com/inforberi/auth-service/internal/delivery/http/handlers/email"
	sessionHandler "github.com/inforberi/auth-service/internal/delivery/http/handlers/session"
	infraPg "github.com/inforberi/auth-service/internal/infra/postgres"
	"github.com/inforberi/auth-service/internal/infra/redis"
	infraRedis "github.com/inforberi/auth-service/internal/infra/redis"
	"github.com/inforberi/auth-service/internal/pkg"
	authpg "github.com/inforberi/auth-service/internal/repository/postgres/authpg"
	sessionpg "github.com/inforberi/auth-service/internal/repository/postgres/sessionpg"
	"github.com/inforberi/auth-service/internal/repository/redis/sessionredis"
	"github.com/inforberi/auth-service/internal/service/auth/email"
	"github.com/inforberi/auth-service/internal/service/auth/session"

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
	pgPool, err := infraPg.NewPgPool(ctx, &cfg.Postgres)
	if err != nil {
		log.Error("failed to create pg pool", "err", err)
		return nil, err
	}

	// redis
	redisClient, err := infraRedis.New(ctx, cfg.Redis)
	if err != nil {
		log.Error("failed to create redis db", "err", err)
		return nil, err
	}

	// redis repo
	redisRepo := sessionredis.New(redisClient.Raw())

	// repo
	authRepo := authpg.New(pgPool)
	sessionRepo := sessionpg.New(pgPool)

	// service deps
	clock := pkg.SystemClock{}
	hasher := email.Argon2idHasher{}
	token := pkg.SecureTokenGenerator{}

	// services
	sessionService := session.New(sessionRepo, token, clock, &cfg.Auth, redisRepo)
	emailService := email.New(authRepo, clock, hasher, sessionService)

	// handlers
	emailHandler := emailHandler.New(emailService, log)
	sessionHandler := sessionHandler.New(sessionService, log)

	// router
	router := router.NewRouter(emailHandler, sessionHandler, emailService, &cfg.HTTP)

	return &App{
		Log:         log,
		pgPool:      pgPool,
		Router:      router,
		redisClient: redisClient,
	}, nil
}
