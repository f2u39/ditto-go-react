package mgo

import (
	"context"
	"ditto/lib/format"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

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

func Aggregate(col *mongo.Collection, pipeline []bson.D, T any) error {
	cur, err := col.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Println(err)
		return err
	}

	resultv := reflect.ValueOf(T)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("Result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0

	for cur.Next(context.TODO()) {
		elemp := reflect.New(elemt)
		err := cur.Decode(&elemp)
		if err != nil {
			log.Println(err)
			return err
		}

		slicev = reflect.Append(slicev, elemp.Elem())
		slicev = slicev.Slice(0, slicev.Cap())
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return err
	}

	resultv.Elem().Set(slicev.Slice(0, i))

	cur.Close(context.TODO())
	return nil
}

func Count(col *mongo.Collection, filter any) (int64, error) {
	return col.CountDocuments(context.TODO(), filter)
}

func DeleteID(col *mongo.Collection, id any) error {
	var objID primitive.ObjectID

	switch id.(type) {
	case string:
		objID = format.ToObjID(fmt.Sprintf("%v", id))
	case primitive.ObjectID:
		objID = id.(primitive.ObjectID)
	default:
		return nil
	}

	_, err := col.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func Insert(col *mongo.Collection, T any) error {
	_, err := col.InsertOne(context.TODO(), T)
	if err != nil {
		log.Println(err)
	}
	return err
}

func FindOne(col *mongo.Collection, filter any, T any) error {
	return col.FindOne(context.TODO(), filter).Decode(&T)
}

func FindID(col *mongo.Collection, id any, T any) error {
	var objID primitive.ObjectID

	switch id.(type) {
	case string:
		objID = format.ToObjID(fmt.Sprintf("%v", id))
	case primitive.ObjectID:
		objID = id.(primitive.ObjectID)
	case nil:
		return fmt.Errorf("bad id format")
	default:
		return fmt.Errorf("bad id format")
	}

	return col.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&T)
}

func FindMany(col *mongo.Collection, T any, filter any, sorts ...any) error {
	options := options.Find()

	if len(sorts) > 0 {
		options.SetSort(sorts)
	}

	cur, err := col.Find(context.TODO(), filter, options)
	if err != nil {
		log.Println(err)
		return err
	}

	resultv := reflect.ValueOf(T)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("Result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0

	for cur.Next(context.TODO()) {
		elemp := reflect.New(elemt)
		err := cur.Decode(&elemp)
		if err != nil {
			log.Println(err)
			return err
		}

		slicev = reflect.Append(slicev, elemp.Elem())
		slicev = slicev.Slice(0, slicev.Cap())
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return err
	}

	resultv.Elem().Set(slicev.Slice(0, i))

	cur.Close(context.TODO())
	return nil
}

func FindPage(col *mongo.Collection, T any, filter any, page, limit int, sorts ...string) (int, error) {
	_page := int64(page)
	_limit := int64(limit)

	opts := options.Find()

	if len(sorts) > 0 {
		opts.SetSort(sorts)
	}

	if _page <= 0 {
		_page = 1
	}

	cnt, err := Count(col, filter)
	if err != nil {
		return 0, err
	}

	totalPages := cnt / _limit
	if cnt%_limit > 0 {
		totalPages += 1
	}
	if totalPages == 0 {
		totalPages = 1
	}

	opts.SetSkip(int64((_page - 1) * _limit)).SetLimit(_limit)

	cur, err := col.Find(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	resultv := reflect.ValueOf(T)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("Result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0

	for cur.Next(context.TODO()) {
		elemp := reflect.New(elemt)
		err := cur.Decode(&elemp)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		slicev = reflect.Append(slicev, elemp.Elem())
		slicev = slicev.Slice(0, slicev.Cap())
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return 0, err
	}

	resultv.Elem().Set(slicev.Slice(0, i))

	cur.Close(context.TODO())

	return int(totalPages), nil
}

func Update(col *mongo.Collection, id any, upd any) error {
	var objID primitive.ObjectID

	switch id.(type) {
	case string:
		objID = format.ToObjID(fmt.Sprintf("%v", id))
	case primitive.ObjectID:
		objID = id.(primitive.ObjectID)
	default:
		return nil
	}

	_, err := col.UpdateOne(context.TODO(), bson.M{"_id": objID}, upd)
	return err
}
