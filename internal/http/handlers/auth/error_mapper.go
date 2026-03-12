package auth

import (
	"errors"
	"net/http"

	"github.com/inforberi/auth-service/internal/service/auth"
)

func mapAuthError(err error) (status int, code string, message string, ok bool) {
	switch {
	case errors.Is(err, auth.ErrEmailTaken):
		return http.StatusConflict, "email_taken", err.Error(), true

	case errors.Is(err, auth.ErrEmptyEmail),
		errors.Is(err, auth.ErrInvalidEmail),
		errors.Is(err, auth.ErrPasswordTooShort),
		errors.Is(err, auth.ErrPasswordTooLong),
		errors.Is(err, auth.ErrPasswordNoLetter),
		errors.Is(err, auth.ErrPasswordNoDigit):
		return http.StatusBadRequest, "validation_error", err.Error(), true

	case errors.Is(err, auth.ErrInvalidCredentials):
		return http.StatusUnauthorized, "invalid_credentials", err.Error(), true

	case errors.Is(err, auth.ErrUserDisabled):
		return http.StatusForbidden, "user_disabled", err.Error(), true

	case errors.Is(err, auth.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized", "invalid session", true

	default:
		return 0, "", "", false
	}
}
