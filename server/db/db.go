package db

import (
	"ditto/db/mongo"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

var (
	// MySQL database
	// MySqlDb *sql.DB

	// MongoDB session
	MgoSess *mgo.Session

	// Redis store
	// RedisStore *redis.Store
	RedisCli *redis.Client
)

// Return a copy of session
func Mgo() *mgo.Session {
	return MgoSess.Copy()
}

func Init() {
	mongo.Init()
	// firebase.Init()
}
