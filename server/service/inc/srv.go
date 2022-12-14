package inc

import (
	"ditto/db/mgo"
	"ditto/model/inc"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type incService struct{}

type IncService interface {
	All() []inc.Inc
	IsExists(name string) bool
	ByID(id any) inc.Inc
	ByName(name string) (inc.Inc, error)
	Create(inc inc.Inc) error
	Developers() []inc.Inc
	Publishers() []inc.Inc
	Update(inc inc.Inc) error
}

func NewIncService() IncService {
	return &incService{}
}

func (s *incService) All() []inc.Inc {
	var incs []inc.Inc
	mgo.FindMany(mgo.Incs, &incs, bson.D{}, bson.D{})
	return incs
}

func (s *incService) IsExists(name string) bool {
	var inc inc.Inc

	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result := mgo.FindOne(mgo.Incs, filter)
	err := result.Decode(&inc)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false
		}
	}

	return true
}

func (s *incService) ByID(id any) inc.Inc {
	var inc inc.Inc
	result := mgo.FindID(mgo.Incs, id)
	result.Decode(&inc)
	return inc
}

func (s *incService) ByName(name string) (inc.Inc, error) {
	var inc inc.Inc

	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result := mgo.FindOne(mgo.Incs, filter)
	err := result.Decode(&inc)
	return inc, err
}

func (s *incService) Create(inc inc.Inc) error {
	inc.ID = primitive.NewObjectIDFromTimestamp(time.Now())
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
