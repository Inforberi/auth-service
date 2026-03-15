package sessionredis

import (
	"context"
	"encoding/json"
	"time"
)

func (s *Store) SetSession(ctx context.Context, tokenHash []byte, snap SessionSnapshot,
	ttl time.Duration) error {
	if ttl <= 0 {
		return nil
	}

	key := sessionKey(tokenHash)

	data, err := json.Marshal(snap)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, key, data, ttl).Err()
}
