package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/http/handlers/auth"
	"github.com/inforberi/auth-service/internal/http/handlers/health"
	"github.com/inforberi/auth-service/internal/http/handlers/session"
)

func NewRouter(authHandler *auth.AuthHandler, sessionHandler *session.SessionHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", health.Health)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register/email", authHandler.RegisterEmail)
			r.Post("/login/email", authHandler.LoginEmail)

			r.Post("/logout", sessionHandler.Logout)
			r.Post("/logout-all", sessionHandler.LogoutAll)

			r.Get("/me", authHandler.Me)
		})
	})

	return r
}
