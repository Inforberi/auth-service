package email

import (
	"errors"
	"net/http"

	"github.com/inforberi/auth-service/internal/service/auth/email"
)

func mapAuthError(err error) (status int, code string, message string, ok bool) {
	switch {
	case errors.Is(err, email.ErrEmailTaken):
		return http.StatusConflict, "email_taken", err.Error(), true

	case errors.Is(err, email.ErrEmptyEmail),
		errors.Is(err, email.ErrInvalidEmail),
		errors.Is(err, email.ErrPasswordTooShort),
		errors.Is(err, email.ErrPasswordTooLong),
		errors.Is(err, email.ErrPasswordNoLetter),
		errors.Is(err, email.ErrPasswordNoDigit):
		return http.StatusBadRequest, "validation_error", err.Error(), true

	case errors.Is(err, email.ErrInvalidCredentials):
		return http.StatusUnauthorized, "invalid_credentials", err.Error(), true

	case errors.Is(err, email.ErrUserDisabled):
		return http.StatusForbidden, "user_disabled", err.Error(), true

	case errors.Is(err, email.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized", "invalid session", true

	default:
		return 0, "", "", false
	}
}
