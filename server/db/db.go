package db

import (
	"ditto/db/mongo"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Redis store
	// RedisStore *redis.Store
	RedisCli *redis.Client
)

func Init() {
	mongo.Init()
}
