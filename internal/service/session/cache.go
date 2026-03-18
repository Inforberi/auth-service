package session

import (
	"context"
	"time"

	"github.com/inforberi/auth-service/internal/model/sessionmodel"
)

func (s *SessionService) setSessionCache(
	ctx context.Context,
	tokenHash []byte,
	userID string,
	sessionID string,
	sessionVersion int,
	actualSessionVersion int,
	expiresAt time.Time,
	revokedAt *time.Time,
	disabledAt *time.Time,
) {
	if s.cache == nil {
		return
	}

	ttl := time.Until(expiresAt)
	if ttl > 0 {
		_ = s.cache.SetSession(ctx, tokenHash, sessionmodel.CacheSession{
			SessionID:      sessionID,
			UserID:         userID,
			SessionVersion: sessionVersion,
			ExpiresAtUnix:  expiresAt.UTC().Unix(),
			Revoked:        revokedAt != nil,
			UserDisabled:   disabledAt != nil,
		}, ttl)
	}

	_ = s.cache.SetUserSessionVersion(ctx, userID, actualSessionVersion)
}

func (s *SessionService) getSessionFromCache(
	ctx context.Context,
	tokenHash []byte,
	now time.Time,
) (GetSessionResult, bool, error) {
	if s.cache == nil {
		return GetSessionResult{}, false, nil
	}

	cached, found, err := s.cache.GetSession(ctx, tokenHash)
	if err != nil || !found {
		return GetSessionResult{}, false, nil
	}

	if cached.Revoked {
		return GetSessionResult{}, false, ErrSessionIsRevoked
	}

	expiresAt := time.Unix(cached.ExpiresAtUnix, 0).UTC()
	if !expiresAt.After(now) {
		return GetSessionResult{}, false, ErrSessionIsExpired
	}

	if cached.UserDisabled {
		return GetSessionResult{}, false, ErrUserIsDisabled
	}

	actualSessionVersion, foundVersion, err := s.cache.GetUserSessionVersion(ctx, cached.UserID)
	if err != nil || !foundVersion {
		return GetSessionResult{}, false, nil
	}

	if actualSessionVersion != cached.SessionVersion {
		return GetSessionResult{}, false, ErrSessionVersionMismatch
	}

	return GetSessionResult{
		SessionID:      cached.SessionID,
		UserID:         cached.UserID,
		SessionVersion: cached.SessionVersion,
	}, true, nil
}

func (s *SessionService) deleteSessionCache(ctx context.Context, tokenHash []byte) {
	if s.cache == nil {
		return
	}

	_ = s.cache.DeleteSession(ctx, tokenHash)
}
