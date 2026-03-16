package sessionredis

import (
	"context"
	"encoding/json"

	"github.com/inforberi/auth-service/internal/model/sessionmodel"
	"github.com/redis/go-redis/v9"
)

func (s *Store) GetSession(ctx context.Context, tokenHash []byte) (sessionmodel.CacheSession, bool, error) {
	if s == nil || s.rdb == nil {
		return sessionmodel.CacheSession{}, false, nil
	}

	key := sessionKey(tokenHash)

	data, err := s.rdb.Get(ctx, key).Bytes()
	if err != nil {

		if err == redis.Nil {
			return sessionmodel.CacheSession{}, false, nil
		}

		return sessionmodel.CacheSession{}, false, err
	}

	var snap sessionmodel.CacheSession

	if err := json.Unmarshal(data, &snap); err != nil {
		return sessionmodel.CacheSession{}, false, err
	}

	return snap, true, nil
}
