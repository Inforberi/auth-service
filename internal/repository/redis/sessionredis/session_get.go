package sessionredis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func (s *Store) GetSession(ctx context.Context, tokenHash []byte) (SessionSnapshot, bool, error) {
	key := sessionKey(tokenHash)

	data, err := s.rdb.Get(ctx, key).Bytes()
	if err != nil {

		if err == redis.Nil {
			return SessionSnapshot{}, false, nil
		}

		return SessionSnapshot{}, false, err
	}

	var snap SessionSnapshot

	if err := json.Unmarshal(data, &snap); err != nil {
		return SessionSnapshot{}, false, err
	}

	return snap, true, nil
}
