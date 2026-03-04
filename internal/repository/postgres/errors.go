package postgres

import "errors"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrEmailTaken      = errors.New("email already taken")
)
