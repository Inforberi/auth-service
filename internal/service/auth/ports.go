package auth

import (
	"context"
	"time"
)

type Auth interface {
	CreateUserWithEmailPassword(ctx context.Context, email, emailNorm, passwordHash string, now time.Time) (string, error)
}

type Clock interface {
	Now() time.Time
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(storedHash, password string) bool
}
