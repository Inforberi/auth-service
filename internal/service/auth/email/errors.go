package email

import (
	"errors"

	"github.com/inforberi/auth-service/internal/service/auth/session"
)

var (
	// Email
	ErrEmptyEmail   = errors.New("Email is empty")
	ErrInvalidEmail = errors.New("invalid email")
	ErrEmailTaken   = errors.New("email already taken")

	// Password
	ErrPasswordTooShort = errors.New("password too short")
	ErrPasswordTooLong  = errors.New("password too long")
	ErrPasswordNoLetter = errors.New("password must contain at least one letter")
	ErrPasswordNoDigit  = errors.New("password must contain at least one digit")
	ErrHashFormat       = errors.New("invalid password hash format")

	// Auth
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserDisabled       = errors.New("user is disabled")
	ErrLogin              = errors.New("login failed")
	ErrUnauthorized       = errors.New("unauthorized")

	// Provider
	ErrProviderNotEnabled = errors.New("Provider is not enabled")

	// Common
	ErrRegister = errors.New("register failed")
)

func isUnauthorizedSessionError(err error) bool {
	switch {
	case errors.Is(err, session.ErrSessionNotFound),
		errors.Is(err, session.ErrSessionIsRevoked),
		errors.Is(err, session.ErrSessionIsExpired),
		errors.Is(err, session.ErrSessionVersionMismatch),
		errors.Is(err, session.ErrUserIsDisabled):
		return true
	default:
		return false
	}
}

type repoEmailTaken interface {
	EmailTaken() bool
}

func isRepoEmailTaken(err error) bool {
	var marker repoEmailTaken
	return errors.As(err, &marker) && marker.EmailTaken()
}
