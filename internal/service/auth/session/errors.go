package session

import "errors"

type repoSessionNotFound interface {
	SessionNotFound() bool
}

type repoUserNotFound interface {
	UserNotFound() bool
}

func isRepoSessionNotFound(err error) bool {
	var marker repoSessionNotFound
	return errors.As(err, &marker) && marker.SessionNotFound()
}

func isRepoUserNotFound(err error) bool {
	var marker repoUserNotFound
	return errors.As(err, &marker) && marker.UserNotFound()
}

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
	ErrUserNotFound           = errors.New("User not found")
	ErrLogoutAll              = errors.New("logout all failed")

	// redis
	ErrCacheSync = errors.New("cache sync failed")
)
