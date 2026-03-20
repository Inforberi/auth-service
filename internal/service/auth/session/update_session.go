package session

import (
	"context"
	"time"

	"github.com/inforberi/auth-service/internal/model/sessionmodel"
)

func (s *SessionService) UpdateSessionActivity(
	ctx context.Context,
	sessionID string,
	tokenHash []byte,
) error {
	now := s.clock.Now().UTC()

	if s.cache != nil {
		cached, found, err := s.cache.GetSession(ctx, tokenHash)
		if err == nil && found {
			expiresAt := time.Unix(cached.ExpiresAtUnix, 0).UTC()

			// refresh <= 24ч
			refreshBeforeExpiry := 24 * time.Hour
			if expiresAt.After(now.Add(refreshBeforeExpiry)) {
				return nil
			}
		}
	}

	newExpiresAt := now.Add(s.sessionTTL)
	lastSeenThreshold := now.Add(-s.auth.UpdateInterval)
	refreshBefore := now.Add(s.auth.RefreshBeforeExpiry)

	// atomic Postgres
	updated, err := s.repo.UpdateSessionActivity(
		ctx,
		sessionID,
		now,
		newExpiresAt,
		lastSeenThreshold,
		refreshBefore,
	)
	if err != nil {
		return ErrUpdateSessionLastSeen
	}

	if !updated {
		return nil
	}

	// if db updated -> update redis
	if s.cache != nil {
		cached, found, err := s.cache.GetSession(ctx, tokenHash)
		if err == nil && found {
			ttl := time.Until(newExpiresAt)
			if ttl > 0 {
				_ = s.cache.SetSession(ctx, tokenHash, sessionmodel.CacheSession{
					SessionID:      cached.SessionID,
					UserID:         cached.UserID,
					SessionVersion: cached.SessionVersion,
					ExpiresAtUnix:  newExpiresAt.UTC().Unix(),
					Revoked:        cached.Revoked,
					UserDisabled:   cached.UserDisabled,
				}, ttl)
			}
		}
	}

	return nil
}
