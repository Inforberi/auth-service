package middleware

import (
	"net/http"
	"strings"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
)

func RequireJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead || r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			helpers.WriteError(w, http.StatusUnsupportedMediaType, "unsupported_media_type", "content type must be application/json")
			return
		}

		mediaType := strings.TrimSpace(strings.Split(contentType, ";")[0])
		if mediaType != "application/json" {
			helpers.WriteError(w, http.StatusUnsupportedMediaType, "unsupported_media_type", "content type must be application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}
