package inc

import (
	"ditto/db/mgo"
	"ditto/model/inc"
	"ditto/service/base"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncService interface {
	All() []inc.Inc
	ByID(id any) inc.Inc
	Create(inc inc.Inc) error
	Developers() []inc.Inc
	Publishers() []inc.Inc
	Update(inc inc.Inc) error
}

func NewIncService() IncService {
	return &incService{Base: base.NewBaseRepo()}
}

type incService struct {
	Base base.BaseRepo
}

func (s *incService) All() []inc.Inc {
	var incs []inc.Inc
	mgo.FindMany(mgo.Incs, &incs, bson.D{}, bson.D{})
	return incs
}

func (s *incService) ByID(id any) inc.Inc {
	var inc inc.Inc
	result, err := mgo.FindID(mgo.Incs, id)
	if err != nil {
		return inc
	}
	result.Decode(&inc)
	return inc
}

func (s *incService) Create(inc inc.Inc) error {
	inc.CreatedAt = time.Now()
	inc.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Incs, inc)
}

func (s *incService) Developers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.D{primitive.E{Key: "is_developer", Value: 1}}
	srt := bson.D{primitive.E{Key: "name", Value: 1}}
	mgo.FindMany(mgo.Incs, &incs, qry, srt)
	return incs
}

func (s *incService) Publishers() []inc.Inc {
	var incs []inc.Inc
	qry := bson.D{primitive.E{Key: "is_publisher", Value: 1}}
	srt := bson.D{primitive.E{Key: "name", Value: 1}}
	mgo.FindMany(mgo.Incs, &incs, qry, srt)
	return incs
}

func (s *incService) Update(inc inc.Inc) error {
	inc.CreatedAt = time.Now()
	inc.UpdatedAt = time.Now()
	return mgo.Update(mgo.Incs, inc.ID, inc)
}
