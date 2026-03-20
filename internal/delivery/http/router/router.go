package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/delivery/http/handlers/email"
	"github.com/inforberi/auth-service/internal/delivery/http/handlers/health"
	"github.com/inforberi/auth-service/internal/delivery/http/handlers/session"
	emailService "github.com/inforberi/auth-service/internal/service/auth/email"
)

func NewRouter(emailHandler *email.EmailHandler, sessionHandler *session.SessionHandler, authService *emailService.EmailService, cfg *config.HTTP) http.Handler {
	r := chi.NewRouter()

	applyBaseMiddlewares(r, cfg)

	r.Get("/health", health.Health)

	r.Route("/v1", func(r chi.Router) {
		r.Use(globalRateLimit(cfg))

		mountAuthRoutes(r, emailHandler, sessionHandler, authService, cfg)
	})

	return r
}
