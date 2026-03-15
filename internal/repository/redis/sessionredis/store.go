package sessionredis

import goredis "github.com/redis/go-redis/v9"

type Store struct {
	rdb *goredis.Client
}

func NewStore(rdb *goredis.Client) *Store {
	if rdb == nil {
		return nil
	}

	return &Store{rdb: rdb}
}
