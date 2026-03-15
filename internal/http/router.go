package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/http/handlers/auth"
	"github.com/inforberi/auth-service/internal/http/handlers/health"
	"github.com/inforberi/auth-service/internal/http/handlers/session"
	"github.com/inforberi/auth-service/internal/http/middleware"
	authService "github.com/inforberi/auth-service/internal/service/auth"
)

func NewRouter(authHandler *auth.AuthHandler, sessionHandler *session.SessionHandler, authService *authService.AuthService, cfg *config.HTTP) http.Handler {
	r := chi.NewRouter()

	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Timeout(60 * time.Second))
	r.Use(middleware.SecurityHeaders(middleware.SecurityHeadersConfig{
		EnableHSTS: cfg.SecurityHeaders.EnableHSTS,
		HSTSMaxAge: cfg.SecurityHeaders.HSTSMaxAge,
	}))
	r.Use(middleware.CORS(middleware.CORSConfig{
		AllowedOrigins:   cfg.CORS.AllowedOrigins,
		AllowedMethods:   cfg.CORS.AllowedMethods,
		AllowedHeaders:   cfg.CORS.AllowedHeaders,
		ExposedHeaders:   cfg.CORS.ExposedHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
		MaxAge:           cfg.CORS.MaxAge,
	}))

	r.Get("/health", health.Health)

	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.RateLimitByIP(middleware.RateLimitConfig{
			Requests: cfg.RateLimit.GlobalPerMinute,
			Window:   time.Minute,
		}))

		r.Route("/auth", func(r chi.Router) {
			r.With(middleware.RateLimitByIP(middleware.RateLimitConfig{
				Requests: cfg.RateLimit.RegisterIPRequests,
				Window:   cfg.RateLimit.RegisterIPWindow,
			})).Post("/register/email", authHandler.RegisterEmail)

			r.With(middleware.RateLimitByIP(middleware.RateLimitConfig{
				Requests: cfg.RateLimit.LoginIPRequests,
				Window:   cfg.RateLimit.LoginIPWindow,
			})).Post("/login/email", authHandler.LoginEmail)

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
