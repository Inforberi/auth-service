package session

import (
	"context"
	"time"

	"github.com/inforberi/auth-service/internal/model/sessionmodel"
)

type SessionRepo interface {
	// sessions
	CreateSession(ctx context.Context, userID string, sessionVersion int, tokenHash []byte, now, expiresAt time.Time, ip, ua *string, deviceID *string) (sessionID string, err error)
	GetSessionByTokenHash(ctx context.Context, tokenHash []byte) (
		sessionID string,
		userID string,
		sessionVersion int,
		expiresAt time.Time,
		revokedAt *time.Time,
		actualSessionVersion int,
		disabledAt *time.Time,
		found bool,
		err error,
	)
	UpdateSessionActivity(ctx context.Context, sessionID string, now time.Time, expiresAt time.Time, threshold time.Time, refreshBefore time.Time) (bool, error)

	RevokeSession(ctx context.Context, sessionID string, now time.Time) error
	IncrementUserSessionVersion(ctx context.Context, userID string, now time.Time) (int, error)
}

type SessionCache interface {
	GetSession(ctx context.Context, tokenHash []byte) (sessionmodel.CacheSession, bool, error)

	SetSession(
		ctx context.Context,
		tokenHash []byte,
		snap sessionmodel.CacheSession,
		ttl time.Duration,
	) error

	DeleteSession(ctx context.Context, tokenHash []byte) error

	GetUserSessionVersion(ctx context.Context, userID string) (int, bool, error)
	SetUserSessionVersion(ctx context.Context, userID string, version int) error

	IsSessionRevoked(ctx context.Context, tokenHash []byte) (bool, error)
	MarkSessionRevoked(ctx context.Context, tokenHash []byte, ttl time.Duration) error
}

type Clock interface {
	Now() time.Time
}

type TokenGenerator interface {
	New() (string, error)
}
