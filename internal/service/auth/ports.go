package auth

import (
	"context"
	"time"
)

type Auth interface {
	CreateUser(ctx context.Context, now time.Time) (userID string, err error)
}

type Clock interface {
	Now() time.Time
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(storedHash, password string) bool
}
