package sessionredis

import "context"

func (s *Store) DeleteSession(ctx context.Context, tokenHash []byte) error {
	if s == nil || s.rdb == nil {
		return nil
	}

	key := sessionKey(tokenHash)

	return s.rdb.Del(ctx, key).Err()
}
