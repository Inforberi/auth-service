package session

import "errors"

var (
	ErrCreateSession = errors.New("create session failed")
	ErrGenerateToken = errors.New("generate token failed")

	// session
	ErrGetSession             = errors.New("Error get session")
	ErrSessionNotFound        = errors.New("Session not found")
	ErrSessionIsRevoked       = errors.New("Session is revoked")
	ErrSessionIsExpired       = errors.New("Session is expired")
	ErrSessionVersionMismatch = errors.New("Session version mismatch")
	ErrUserIsDisabled         = errors.New("user is disabled")
	ErrUpdateSessionLastSeen  = errors.New("Error update session")
	ErrRevokeSession          = errors.New("Error revoke session")
)
