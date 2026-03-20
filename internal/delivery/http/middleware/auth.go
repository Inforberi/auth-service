package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/service/auth/email"
)

func Auth(authService *email.EmailService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := helpers.ReadSessionCookie(r)
			if err != nil {
				helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing session")
				return
			}

			authInfo, err := authService.Me(r.Context(), token)
			if err != nil {
				if errors.Is(err, email.ErrUnauthorized) {
					helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "invalid session")
					return
				}

				helpers.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
				return
			}

			ctx := context.WithValue(r.Context(), authContextKey, AuthContext{
				UserID:         authInfo.UserID,
				SessionID:      authInfo.SessionID,
				SessionVersion: authInfo.SessionVersion,
				TokenHash:      authInfo.TokenHash,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
