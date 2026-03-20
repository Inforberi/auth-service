package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/http/handlers/auth"
	"github.com/inforberi/auth-service/internal/http/handlers/session"
	"github.com/inforberi/auth-service/internal/http/middleware"
	"github.com/inforberi/auth-service/internal/service/auth/email"
)

func mountAuthRoutes(
	r chi.Router,
	authHandler *auth.AuthHandler,
	sessionHandler *session.SessionHandler,
	emailService *email.EmailService,
	cfg *config.HTTP,
) {
	r.Route("/auth", func(r chi.Router) {
		mountPublicAuthRoutes(r, authHandler, cfg)
		mountProtectedAuthRoutes(r, authHandler, sessionHandler, emailService, cfg)
	})
}

func mountPublicAuthRoutes(
	r chi.Router,
	authHandler *auth.AuthHandler,
	cfg *config.HTTP,
) {
	r.With(registerRouteMiddlewares(cfg)...).
		Post("/register/email", authHandler.RegisterEmail)

	r.With(loginRouteMiddlewares(cfg)...).
		Post("/login/email", authHandler.LoginEmail)
}

func mountProtectedAuthRoutes(
	r chi.Router,
	authHandler *auth.AuthHandler,
	sessionHandler *session.SessionHandler,
	authService *email.EmailService,
	cfg *config.HTTP,
) {
	csrf := csrfMiddleware(cfg)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authService))
		r.Use(middleware.NoStore)

		r.Get("/me", authHandler.Me)
		r.With(csrf).Post("/logout", sessionHandler.Logout)
		r.With(csrf).Post("/logout-all", sessionHandler.LogoutAll)
	})
}
