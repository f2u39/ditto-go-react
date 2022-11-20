package word

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Word struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Word    string             `json:"word" bson:"word"`
	Example string             `json:"example" bson:"example"`
	Meaning string             `json:"meaning" bson:"meaning"`
	// Date      string        `json:"date" bson:"date"`
	IsChecked int       `json:"is_checked" bson:"is_checked"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
