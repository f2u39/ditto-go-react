package inc

import (
	"ditto/db/mgo"
	"ditto/model/inc"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type incRepo struct{}

type IncRepo interface {
	All() []inc.Inc
	Developers() []inc.Inc
	Publishers() []inc.Inc
	ByID(id string) inc.Inc
	Create(inc inc.Inc) error
}

func NewIncRepo() IncRepo {
	return &incRepo{}
}

func (*incRepo) All() []inc.Inc {
	var incs []inc.Inc
	mgo.FindMany(mgo.Incs, &incs, bson.D{})
	return incs
}

func (*incRepo) ByID(id string) inc.Inc {
	inc := inc.Inc{}
	mgo.FindID(mgo.Incs, id, &inc)
	return inc
}

func (*incRepo) Developers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.D{primitive.E{Key: "is_developer", Value: 1}}
	srt := bson.D{primitive.E{Key: "name", Value: 1}}
	mgo.FindMany(mgo.Incs, &incs, qry, srt)
	return incs
}

func (*incRepo) Publishers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.D{primitive.E{Key: "is_publisher", Value: 1}}
	srt := bson.D{primitive.E{Key: "name", Value: 1}}
	mgo.FindMany(mgo.Incs, &incs, qry, srt)
	return incs
}

func (*incRepo) Create(inc inc.Inc) error {
	inc.CreatedAt = time.Now()
	inc.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Incs, inc)
}
