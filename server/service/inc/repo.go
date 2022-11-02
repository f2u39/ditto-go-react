package inc

import (
	"ditto/db/mgo"
	"ditto/model/inc"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type incRepo struct{}

type IncRepo interface {
	All() []inc.Inc
	Developers() []inc.Inc
	Publishers() []inc.Inc
	ByID(id string) inc.Inc
	Create(inc inc.Inc) bool
}

func NewIncRepo() IncRepo {
	return &incRepo{}
}

func (*incRepo) All() []inc.Inc {
	var incs []inc.Inc
	mgo.FindMany(mgo.Incs, &incs, bson.M{})
	return incs
}

func (*incRepo) ByID(id string) inc.Inc {
	inc := inc.Inc{}
	mgo.FindID(mgo.Incs, id, &inc)
	return inc
}

func (*incRepo) Developers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.M{"is_developer": 1}
	srt := "name"
	mgo.FindMany(mgo.Incs, &incs, qry, srt)
	return incs
}

func (*incRepo) Publishers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.M{"is_publisher": 1}
	srt := "name"
	mgo.FindMany(mgo.Incs, &incs, qry, srt)
	return incs
}

func (*incRepo) Create(inc inc.Inc) bool {
	// inc.ID = bson.NewObjectId()
	inc.CreatedAt = time.Now()
	inc.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Incs, inc)
}
