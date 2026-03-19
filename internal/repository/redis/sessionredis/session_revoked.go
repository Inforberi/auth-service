package sessionredis

import (
	"context"
	"time"
)

func (s *Store) IsSessionRevoked(ctx context.Context, tokenHash []byte) (bool, error) {
	if s == nil || s.rdb == nil {
		return false, nil
	}

	key := revokedSessionKey(tokenHash)

	n, err := s.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}

func (s *Store) MarkSessionRevoked(ctx context.Context, tokenHash []byte, ttl time.Duration) error {
	if s == nil || s.rdb == nil {
		return nil
	}

	if ttl <= 0 {
		return nil
	}

	key := revokedSessionKey(tokenHash)

	return s.rdb.Set(ctx, key, "1", ttl).Err()
}
