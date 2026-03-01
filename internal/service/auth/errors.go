package auth

import "errors"

var (
	ErrHashFormat = errors.New("invalid password hash format")
	ErrEmptyEmail = errors.New("Email is empty")
)
