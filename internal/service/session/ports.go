package session

import (
	"context"
	"time"
)

type SessionRepo interface {
	// sessions
	CreateSession(ctx context.Context, userID string, sessionVersion int, tokenHash []byte, now, expiresAt time.Time, ip, ua *string, deviceID *string) (sessionID string, err error)
	GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (
		GetRepoResult, error)
	UpdateSessionLastSeen(ctx context.Context, sessionID string, now time.Time) error

	RevokeSession(ctx context.Context, sessionID string, now time.Time) error
	IncrementUserSessionVersion(ctx context.Context, userID string, now time.Time) error
}

type Clock interface {
	Now() time.Time
}

type TokenGenerator interface {
	New() (string, error)
}
