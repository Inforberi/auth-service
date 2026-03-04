package auth

import (
	"context"
	"time"
)

type Auth interface {
	// register
	CreateUserWithEmailPassword(ctx context.Context, email, emailNorm, passwordHash string, now time.Time) (string, error)

	// login: найти пользователя + пароль
	GetUserByEmail(ctx context.Context, emailNorm string) (userID string, passwordHash string, sessionVersion int, disabledAt *time.Time, found bool, err error)

	// sessions
	CreateSession(ctx context.Context, userID string, sessionVersion int, tokenHash []byte, now, expiresAt time.Time, ip, ua *string) (sessionID string, err error)
	// GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (sessionID string, userID string, sessionVersion int, expiresAt time.Time, revokedAt *time.Time, err error)
	// RevokeSession(ctx context.Context, sessionID string, now time.Time) error
}

type Clock interface {
	Now() time.Time
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(storedHash, password string) bool
}
