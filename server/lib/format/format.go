package format

import "gopkg.in/mgo.v2/bson"

// Transform ObjectIdHex to ObjectId
func ToObjId(objIdHex string) bson.ObjectId {
	if !bson.IsObjectIdHex(objIdHex) {
		return ""
	}
	return bson.ObjectIdHex(objIdHex)
}
