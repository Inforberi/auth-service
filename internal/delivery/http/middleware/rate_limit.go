package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/httprate"
)

type RateLimitConfig struct {
	Requests int
	Window   time.Duration
}

func RateLimitByIP(cfg RateLimitConfig) func(http.Handler) http.Handler {
	requests := cfg.Requests
	if requests <= 0 {
		requests = 5
	}

	window := cfg.Window
	if window <= 0 {
		window = time.Minute
	}

	return httprate.Limit(
		requests,
		window,
		httprate.WithKeyFuncs(httprate.KeyByIP),
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Retry-After", strconv.Itoa(int(window.Seconds())))
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"code":"rate_limited","message":"too many requests"}`))
		}),
	)
}
