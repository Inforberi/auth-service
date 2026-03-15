package middleware

import (
	"net/http"
	"strings"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
)

type CSRFConfig struct {
	AllowedOrigins []string
}

func CSRF(cfg CSRFConfig) func(http.Handler) http.Handler {
	allowed := make(map[string]struct{}, len(cfg.AllowedOrigins))
	for _, origin := range cfg.AllowedOrigins {
		origin = strings.TrimSpace(origin)
		if origin == "" {
			continue
		}
		allowed[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
			default:
				next.ServeHTTP(w, r)
				return
			}

			origin := r.Header.Get("Origin")
			if origin == "" {
				helpers.WriteError(w, http.StatusForbidden, "invalid_origin", "missing origin")
				return
			}

			if _, ok := allowed[origin]; !ok {
				helpers.WriteError(w, http.StatusForbidden, "invalid_origin", "origin not allowed")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
