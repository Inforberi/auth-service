package sessionredis

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func (s *Store) GetUserSessionVersion(ctx context.Context, userID string) (int, bool, error) {
	if s == nil || s.rdb == nil {
		return 0, false, nil
	}

	key := userSessionVersionKey(userID)

	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, false, nil
		}
		return 0, false, err
	}

	version, err := strconv.Atoi(val)
	if err != nil {
		return 0, false, err
	}

	return version, true, nil
}
