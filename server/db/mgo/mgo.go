// Source almost from [mgo](https://github.com/go-mgo/mgo)
package mgo

import (
	"context"
	"ditto/lib/format"
	"fmt"
	"log"
	"reflect"
	"sync"

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

type Iter struct {
	m   sync.Mutex
	cur *mongo.Cursor
	err error
}

func NewIter(cur *mongo.Cursor, err error) *Iter {
	return &Iter{
		cur: cur,
		err: err,
	}
}

func (iter *Iter) Next(T any) bool {
	iter.m.Lock()
	defer iter.m.Unlock()

	if iter.cur.Next(context.TODO()) {
		err := iter.cur.Decode(T)
		if err != nil {
			log.Println(err)
			return false
		}
	} else {
		return false
	}

	// return iter.cur.Err() != nil
	return true
}

func Connect() {
	opts := options.Client().ApplyURI("mongodb://mongo:27017")
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

func Aggregate(col *mongo.Collection, pipe []bson.D, T any) error {
	cur, err := col.Aggregate(context.TODO(), pipe)
	iter := NewIter(cur, err)
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

	for {
		if slicev.Len() == i {
			elemp := reflect.New(elemt)
			if !iter.Next(elemp.Interface()) {
				break
			}
			slicev = reflect.Append(slicev, elemp.Elem())
			slicev = slicev.Slice(0, slicev.Cap())
		} else {
			if !iter.Next(slicev.Index(i).Addr().Interface()) {
				break
			}
		}
		i++
	}

	resultv.Elem().Set(slicev.Slice(0, i))
	return cur.Close(context.TODO())
}

func Count(col *mongo.Collection, filter bson.D) (int64, error) {
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

	_, err := col.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objID}})
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

func FindID(col *mongo.Collection, id any) (*mongo.SingleResult, error) {
	var objID primitive.ObjectID

	switch id.(type) {
	case string:
		objID = format.ToObjID(fmt.Sprintf("%v", id))
	case primitive.ObjectID:
		objID = id.(primitive.ObjectID)
	case nil:
		return nil, fmt.Errorf("bad id format")
	default:
		return nil, fmt.Errorf("bad id format")
	}
	return col.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objID}}), nil
}

func FindMany(col *mongo.Collection, T any, filter bson.D, sorts bson.D) error {
	opts := options.Find()

	if len(sorts) > 0 {
		opts.SetSort(sorts)
	}

	cur, err := col.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
		return err
	}
	iter := NewIter(cur, err)

	resultv := reflect.ValueOf(T)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("Result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0

	for {
		if slicev.Len() == i {
			elemp := reflect.New(elemt)
			if !iter.Next(elemp.Interface()) {
				break
			}
			slicev = reflect.Append(slicev, elemp.Elem())
			slicev = slicev.Slice(0, slicev.Cap())
		} else {
			if !iter.Next(slicev.Index(i).Addr().Interface()) {
				break
			}
		}
		i++
	}

	resultv.Elem().Set(slicev.Slice(0, i))
	return cur.Close(context.TODO())
}

func FindPage(col *mongo.Collection, T any, filter bson.D, page, limit int, sorts bson.D) (int, error) {
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

	cur, err := col.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	iter := NewIter(cur, err)

	resultv := reflect.ValueOf(T)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("Result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0

	for {
		if slicev.Len() == i {
			elemp := reflect.New(elemt)
			if !iter.Next(elemp.Interface()) {
				break
			}
			slicev = reflect.Append(slicev, elemp.Elem())
			slicev = slicev.Slice(0, slicev.Cap())
		} else {
			if !iter.Next(slicev.Index(i).Addr().Interface()) {
				break
			}
		}
		i++
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

	_, err := col.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objID}}, upd)
	return err
}
