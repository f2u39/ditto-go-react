package base

import (
	"ditto/db/mgo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type baseRepo struct{}

type BaseRepo interface {
	All(col *mongo.Collection, T []any, srt ...bson.D) error
	ByID(col *mongo.Collection, id any, T any) error
	Create(col *mongo.Collection, T any) bool
	FindMany(col *mongo.Collection, T any, qry bson.M, srt ...bson.D) error
	Delete(col *mongo.Collection, id any) error
	Update(col *mongo.Collection, id any, upd bson.D) error
}

func NewBaseRepo() BaseRepo {
	return &baseRepo{}
}

func (*baseRepo) All(col *mongo.Collection, []any, srt ...string) error {
	return mgo.FindMany(col, T, bson.M{}, srt...)
}

func (*baseRepo) ByID(col *mongo.Collection, id interface{}, T interface{}) error {
	return mgo.FindID(col, id, T)
}

func (*baseRepo) Create(col *mongo.Collection, T interface{}) bool {
	return mgo.Insert(col, T)
}

func (*baseRepo) Delete(col *mongo.Collection, id interface{}) error {
	return mgo.DeleteID(col, id)
}

func (*baseRepo) Update(col *mongo.Collection, id interface{}, T interface{}) error {
	return mgo.Update(col, id, T)
}

func (*baseRepo) FindMany(col *mongo.Collection, T interface{}, qry bson.M, srt ...string) error {
	return mgo.FindMany(col, T, qry, srt...)
}
