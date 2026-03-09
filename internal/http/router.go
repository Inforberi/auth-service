package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inforberi/auth-service/internal/http/handlers/auth"
)

func NewRouter(authHandler *auth.AuthHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register/email", authHandler.RegisterEmail)
		})
	})

	return r
}
