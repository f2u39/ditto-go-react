package inc

import (
	"ditto/db/mongo"
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
	mongo.FindMany(mongo.Incs, &incs, bson.M{})
	return incs
}

func (*incRepo) ByID(id string) inc.Inc {
	inc := inc.Inc{}
	mongo.FindByID(mongo.Incs, id, &inc)
	return inc
}

func (*incRepo) Developers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.M{"is_developer": 1}
	srt := "name"
	mongo.FindMany(mongo.Incs, &incs, qry, srt)
	return incs
}

func (*incRepo) Publishers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.M{"is_publisher": 1}
	srt := "name"
	mongo.FindMany(mongo.Incs, &incs, qry, srt)
	return incs
}

func (*incRepo) Create(inc inc.Inc) bool {
	// inc.ID = bson.NewObjectId()
	inc.CreatedAt = time.Now()
	inc.UpdatedAt = time.Now()
	return mongo.Insert(mongo.Incs, inc)
}
