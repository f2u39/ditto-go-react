package base

import (
	"ditto/db/mongo"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type baseRepo struct{}

type BaseRepo interface {
	All(col *mgo.Collection, T interface{}, srt ...string) error
	ByID(col *mgo.Collection, id interface{}, T interface{}) error
	Create(col *mgo.Collection, T interface{}) bool
	FindMany(col *mgo.Collection, T interface{}, qry bson.M, srt ...string) error
	Delete(col *mgo.Collection, id interface{}) error
	Update(col *mgo.Collection, id interface{}, T interface{}) error
}

func NewBaseRepo() BaseRepo {
	return &baseRepo{}
}

func (*baseRepo) All(col *mgo.Collection, T interface{}, srt ...string) error {
	return mongo.FindMany(col, T, bson.M{}, srt...)
}

func (*baseRepo) ByID(col *mgo.Collection, id interface{}, T interface{}) error {
	return mongo.FindByID(col, id, T)
}

func (*baseRepo) Create(col *mgo.Collection, T interface{}) bool {
	return mongo.Insert(col, T)
}

func (*baseRepo) Delete(col *mgo.Collection, id interface{}) error {
	return mongo.DeleteByID(col, id)
}

func (*baseRepo) Update(col *mgo.Collection, id interface{}, T interface{}) error {
	return mongo.Update(col, id, T)
}

func (*baseRepo) FindMany(col *mgo.Collection, T interface{}, qry bson.M, srt ...string) error {
	return mongo.FindMany(col, T, qry, srt...)
}
