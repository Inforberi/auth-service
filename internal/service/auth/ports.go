package auth

import (
	"context"
	"time"

	"github.com/inforberi/auth-service/internal/service/session"
)

type Auth interface {
	// register
	IsProviderEnabled(ctx context.Context, code string) (bool, error)
	CreateUserWithEmailPassword(ctx context.Context, email, emailNorm, passwordHash string, now time.Time) (userID string, sessionVersion int, err error)

	// login: найти пользователя + пароль
	GetUserByEmail(ctx context.Context, emailNorm string) (userID string, passwordHash string, sessionVersion int, disabledAt *time.Time, found bool, err error)
}

type Clock interface {
	Now() time.Time
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(storedHash, password string) bool
}

type SessionCreator interface {
	CreateSession(
		ctx context.Context,
		userID string,
		sessionVersion int,
		ip, ua, deviceID *string,
	) (session.CreateSessionResult, error)

	GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (session.GetSessionResult, error)
}
