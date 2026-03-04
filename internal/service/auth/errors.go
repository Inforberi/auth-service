package auth

import "errors"

var (
	ErrEmptyEmail       = errors.New("Email is empty")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrPasswordTooShort = errors.New("password too short")
	ErrPasswordTooLong  = errors.New("password too long")
	ErrPasswordNoLetter = errors.New("password must contain at least one letter")
	ErrPasswordNoDigit  = errors.New("password must contain at least one digit")
	ErrHashFormat       = errors.New("invalid password hash format")
)
