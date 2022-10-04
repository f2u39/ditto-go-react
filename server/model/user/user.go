package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
