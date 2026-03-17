package sessionredis

import (
	"context"
)

func (s *Store) SetUserSessionVersion(ctx context.Context, userID string, version int) error {
	if s == nil || s.rdb == nil {
		return nil
	}

	key := userSessionVersionKey(userID)

	return s.rdb.Set(ctx, key, version, 0).Err()
}
