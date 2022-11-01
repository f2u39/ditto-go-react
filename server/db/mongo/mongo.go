package mongo

import (
	"context"
	"ditto/lib/format"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

var (
	Acts    *mongo.Collection
	Animes  *mongo.Collection
	Incs    *mongo.Collection
	Games   *mongo.Collection
	Users   *mongo.Collection
	Todos   *mongo.Collection
	Words   *mongo.Collection
	Studios *mongo.Collection
)

func Connect() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := "ditto"
	Acts = client.Database(db).Collection("act")
	Animes = client.Database(db).Collection("anime")
	Incs = client.Database(db).Collection("inc")
	Games = client.Database(db).Collection("game")
	Users = client.Database(db).Collection("user")
	Todos = client.Database(db).Collection("todo")
	Words = client.Database(db).Collection("word")
	Studios = client.Database(db).Collection("studio")
}

func Count(col *mongo.Collection, filter bson.D) (int64, error) {
	return col.CountDocuments(context.TODO(), filter)
}

func DeleteByID(col *mgo.Collection, id any) error {
	var objID primitive.ObjectID

	switch id.(type) {
	case string:
		objID = format.ObjId(fmt.Sprintf("%v", id))
	case primitive.ObjectID:
		objID = id.(primitive.ObjectID)
	default:
		return nil
	}

	return col.RemoveId(objID)
}

func Insert(col *mongo.Collection, T interface{}) error {
	_, err := col.InsertOne(context.TODO(), T)
	if err != nil {
		log.Println(err)
	}
	return err
}

func FindOne(col *mongo.Collection, filter bson.D, T any) error {
	return col.FindOne(context.TODO(), filter).Decode(&T)
}

func FindByID(col *mongo.Collection, id any, T any) error {
	var objID primitive.ObjectID

	switch id.(type) {
	case string:
		objID = format.ObjId(fmt.Sprintf("%v", id))
	case primitive.ObjectID:
		objID = id.(primitive.ObjectID)
	case nil:
		return fmt.Errorf("bad id format")
	default:
		return fmt.Errorf("bad id format")
	}

	return col.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&T)
}

func FindMany(col *mongo.Collection, T any, filter bson.M, sorts ...string) error {
	if len(sorts) == 0 {
		return col.Find(qry).All(T)
	} else {
		return col.Find(qry).Sort(sorts...).All(T)
	}
}

// Return total pages and error
func FindPage(col *mgo.Collection, T any, qry bson.M, page, limit int, sorts ...string) (int, error) {
	// Check is page or limit is 0
	if page <= 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	cnt, err := col.Find(qry).Count()
	if err != nil {
		return 0, err
	}

	totalPages := cnt / limit
	if cnt%limit > 0 {
		totalPages += 1
	}
	if totalPages == 0 {
		totalPages = 1
	}

	if len(sorts) == 0 {
		return totalPages, col.Find(qry).Skip((page - 1) * limit).Limit(limit).All(T)
	} else {
		return totalPages, col.Find(qry).Sort(sorts...).Skip((page - 1) * limit).Limit(limit).All(T)
	}
}

func Update(col *mgo.Collection, id interface{}, T interface{}) error {
	return col.UpdateId(id, T)
}
