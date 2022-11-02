package inc

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// Inc collection
	col = "inc"
)

// Inc model
type Inc struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	IsDeveloper int                `json:"is_dev" bson:"is_developer"`
	IsPublisher int                `json:"is_pub" bson:"is_publisher"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
