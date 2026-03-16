package sessionredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/inforberi/auth-service/internal/model/sessionmodel"
)

func (s *Store) SetSession(ctx context.Context, tokenHash []byte, snap sessionmodel.CacheSession,
	ttl time.Duration) error {
	if s == nil || s.rdb == nil {
		return nil
	}

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
