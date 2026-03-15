package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/http/handlers/auth"
	"github.com/inforberi/auth-service/internal/http/handlers/health"
	"github.com/inforberi/auth-service/internal/http/handlers/session"
	authService "github.com/inforberi/auth-service/internal/service/auth"
)

func NewRouter(authHandler *auth.AuthHandler, sessionHandler *session.SessionHandler, authService *authService.AuthService, cfg *config.HTTP) http.Handler {
	r := chi.NewRouter()

	applyBaseMiddlewares(r, cfg)

	r.Get("/health", health.Health)

	r.Route("/v1", func(r chi.Router) {
		r.Use(globalRateLimit(cfg))

		mountAuthRoutes(r, authHandler, sessionHandler, authService, cfg)
	})

	return r
}
