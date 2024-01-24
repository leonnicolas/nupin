package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/leonnicolas/nupin/config"
)

var _ Store = &RedisStore{}

func NewRedisStore(cfg *config.Config) *RedisStore {
	if cfg.Redis.UseSentinel {
		return &RedisStore{
			c: redis.NewFailoverClient(cfg.Redis.FailoverOptions),
		}
	}
	return &RedisStore{
		c: redis.NewClient(cfg.Redis.Options),
	}
}

type RedisStore struct {
	c *redis.Client
}

func (m *RedisStore) Exist(key string) (bool, error) {
	e := m.c.Exists(context.TODO(), key)
	return e.Val() == 1, e.Err()
}

func (m *RedisStore) Set(key string) error {
	s := m.c.SetEx(context.TODO(), key, nil, 5*time.Minute)
	return s.Err()
}

func (m *RedisStore) Delete(key string) error {
	d := m.c.Del(context.TODO(), key)
	return d.Err()
}
