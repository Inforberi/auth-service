package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/http/handlers/auth"
	"github.com/inforberi/auth-service/internal/http/handlers/health"
	"github.com/inforberi/auth-service/internal/http/handlers/session"
	"github.com/inforberi/auth-service/internal/http/middleware"
	authService "github.com/inforberi/auth-service/internal/service/auth"
)

func NewRouter(authHandler *auth.AuthHandler, sessionHandler *session.SessionHandler, authService *authService.AuthService) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", health.Health)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register/email", authHandler.RegisterEmail)
			r.Post("/login/email", authHandler.LoginEmail)

			r.Group(func(r chi.Router) {
				r.Use(middleware.Auth(authService))
				r.Get("/me", authHandler.Me)
				r.Post("/logout", sessionHandler.Logout)
				r.Post("/logout-all", sessionHandler.LogoutAll)

				// r.Route("/users", func(r chi.Router) {
				// 	r.Get("/profile", usersHandler.GetProfile)
				// 	r.Patch("/profile", usersHandler.UpdateProfile)
				// })
			})
		})

	})

	return r
}
