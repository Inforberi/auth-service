package middleware

import (
	"net/http"
	"strconv"
)

type SecurityHeadersConfig struct {
	EnableHSTS bool
	HSTSMaxAge int
}

func SecurityHeaders(cfg SecurityHeadersConfig) func(http.Handler) http.Handler {

	var hstsValue string

	if cfg.EnableHSTS {
		maxAge := cfg.HSTSMaxAge
		if maxAge <= 0 {
			maxAge = 31536000
		}

		hstsValue = "max-age=" + strconv.Itoa(maxAge) + "; includeSubDomains"
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := w.Header()

			h.Set("X-Content-Type-Options", "nosniff")
			h.Set("X-Frame-Options", "DENY")
			h.Set("Referrer-Policy", "no-referrer")
			h.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

			if hstsValue != "" {
				h.Set("Strict-Transport-Security", hstsValue)
			}

			next.ServeHTTP(w, r)
		})
	}
}
