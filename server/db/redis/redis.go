package redis

import (
	"context"
	"ditto/model/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	Cli *redis.Client
)

func NewRedisClient() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       0,
	})
	return cli
}

func Set(key string, value string, expiration time.Duration) error {
	return Cli.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (interface{}, error) {
	return Cli.Get(ctx, key).Result()
}

func Del(key string) error {
	return Cli.Del(ctx, key).Err()
}
