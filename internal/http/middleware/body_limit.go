package middleware

import "net/http"

func BodyLimit(limit int64) func(http.Handler) http.Handler {
	if limit <= 0 {
		limit = 16 * 1024
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, limit)
			next.ServeHTTP(w, r)
		})
	}
}
