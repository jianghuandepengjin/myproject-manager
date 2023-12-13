package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"test.com/project-user/config"
	"time"
)

var Redis *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.C.InitRedisOptions())
	Redis = &RedisCache{
		rdb: rdb,
	}
}

func (r *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := Redis.rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := Redis.rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}
