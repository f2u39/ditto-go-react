package format

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transform ObjectIdHex to ObjectId
func ToObjID(hex string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(hex)
	return objID
}
