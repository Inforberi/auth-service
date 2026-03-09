package auth

import "errors"

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

	// Provider
	ErrProviderNotEnabled = errors.New("Provider is not enabled")

	// Common
	ErrRegister = errors.New("register failed")
)
