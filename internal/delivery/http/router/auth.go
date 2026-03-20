package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/config"
	emailHandler "github.com/inforberi/auth-service/internal/delivery/http/handlers/email"
	"github.com/inforberi/auth-service/internal/delivery/http/handlers/session"
	"github.com/inforberi/auth-service/internal/delivery/http/middleware"
	"github.com/inforberi/auth-service/internal/service/auth/email"
)

func mountAuthRoutes(
	r chi.Router,
	emailHandler *emailHandler.EmailHandler,
	sessionHandler *session.SessionHandler,
	emailService *email.EmailService,
	cfg *config.HTTP,
) {
	r.Route("/auth", func(r chi.Router) {
		mountPublicAuthRoutes(r, emailHandler, cfg)
		mountProtectedAuthRoutes(r, emailHandler, sessionHandler, emailService, cfg)
	})
}

func mountPublicAuthRoutes(
	r chi.Router,
	emailHandler *emailHandler.EmailHandler,
	cfg *config.HTTP,
) {
	r.With(registerRouteMiddlewares(cfg)...).
		Post("/register/email", emailHandler.RegisterEmail)

	r.With(loginRouteMiddlewares(cfg)...).
		Post("/login/email", emailHandler.LoginEmail)
}

func mountProtectedAuthRoutes(
	r chi.Router,
	authHandler *emailHandler.EmailHandler,
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
