package router

import (
	"net/http"
	"time"

	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/delivery/http/middleware"
)

func globalRateLimit(cfg *config.HTTP) func(http.Handler) http.Handler {
	return middleware.RateLimitByIP(middleware.RateLimitConfig{
		Requests: cfg.RateLimit.GlobalPerMinute,
		Window:   time.Minute,
	})
}

func registerRouteMiddlewares(cfg *config.HTTP) []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.NoStore,
		middleware.RequireJSON,
		middleware.BodyLimit(cfg.BodyLimit.AuthBytes),
		middleware.RateLimitByIP(middleware.RateLimitConfig{
			Requests: cfg.RateLimit.RegisterIPRequests,
			Window:   cfg.RateLimit.RegisterIPWindow,
		}),
	}
}

func loginRouteMiddlewares(cfg *config.HTTP) []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.NoStore,
		middleware.RequireJSON,
		middleware.BodyLimit(cfg.BodyLimit.AuthBytes),
		middleware.RateLimitByIP(middleware.RateLimitConfig{
			Requests: cfg.RateLimit.LoginIPRequests,
			Window:   cfg.RateLimit.LoginIPWindow,
		}),
	}
}

func csrfMiddleware(cfg *config.HTTP) func(http.Handler) http.Handler {
	return middleware.CSRF(middleware.CSRFConfig{
		AllowedOrigins: cfg.CORS.AllowedOrigins,
	})
}
