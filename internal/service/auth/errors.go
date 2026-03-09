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

	// Provider
	ErrProviderNotEnabled = errors.New("Provider is not enabled")

	// Common
	ErrRegister = errors.New("register failed")
)
