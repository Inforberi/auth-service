package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/delivery/http/middleware"
)

func applyBaseMiddlewares(r chi.Router, cfg *config.HTTP) {
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
}
