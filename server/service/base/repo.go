package base

import (
	"ditto/db/mgo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type baseRepo struct{}

type BaseRepo interface {
	All(col *mongo.Collection, T any, srt ...any) error
	ByID(col *mongo.Collection, id any, T any) error
	Create(col *mongo.Collection, T any) error
	FindMany(col *mongo.Collection, T any, filter any, srt ...any) error
	Delete(col *mongo.Collection, id any) error
	Update(col *mongo.Collection, id any, upd any) error
}

func NewBaseRepo() BaseRepo {
	return &baseRepo{}
}

func (*baseRepo) All(col *mongo.Collection, T any, srt ...any) error {
	return mgo.FindMany(col, T, bson.M{}, srt...)
}

func (*baseRepo) ByID(col *mongo.Collection, id any, T any) error {
	return mgo.FindID(col, id, T)
}

func (*baseRepo) Create(col *mongo.Collection, T any) error {
	return mgo.Insert(col, T)
}

func (*baseRepo) Delete(col *mongo.Collection, id any) error {
	return mgo.DeleteID(col, id)
}

func (*baseRepo) FindMany(col *mongo.Collection, T any, filter any, srt ...any) error {
	return mgo.FindMany(col, T, filter, srt...)
}

func (*baseRepo) Update(col *mongo.Collection, id any, upd any) error {
	return mgo.Update(col, id, upd)
}
