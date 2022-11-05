package act

import (
	"ditto/db/mgo"
	"ditto/lib/datetime"
	"ditto/lib/format"
	"ditto/model/act"
	"ditto/service/base"
	"ditto/service/game"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	ByDate(date string) ([]act.Detail, error)
	ByGame(gameId string) ([]act.Act, error)
	ByMonth(date string) ([]act.Detail, error)
	Create(act act.Act) error
	DaySum(ymd string) act.Summary
	Delete(actId string) error
	Duration(ymd string, typ act.Type) int
	MonthSum(yyyymm string) act.Summary
}

type service struct {
	Base base.BaseRepo
}

func NewService() Service {
	return &service{Base: base.NewBaseRepo()}
}

func (s *service) ByDate(date string) ([]act.Detail, error) {
	if len(date) == 0 {
		return nil, fmt.Errorf("unknown date")
	}

	var acts []act.Detail
	match := bson.D{
		primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "date", Value: date}}},
	}

	group := bson.D{
		primitive.E{Key: "$group", Value: bson.D{
			primitive.E{Key: "_id", Value: "$_id"},
			primitive.E{Key: "type", Value: bson.D{primitive.E{Key: "$first", Value: "$type"}}},
			primitive.E{Key: "game_id", Value: bson.D{primitive.E{Key: "$first", Value: "$game_id"}}},
			primitive.E{Key: "duration", Value: bson.D{primitive.E{Key: "$sum", Value: "$duration"}}},
		}},
	}

	sort := bson.D{
		primitive.E{Key: "$sort", Value: bson.D{
			primitive.E{Key: "type", Value: 1},
			primitive.E{Key: "duration", Value: 1},
		}},
	}

	lookup := bson.D{
		primitive.E{Key: "$lookup", Value: bson.D{
			primitive.E{Key: "from", Value: "game"},
			primitive.E{Key: "localField", Value: "game_id"},
			primitive.E{Key: "foreignField", Value: "_id"},
			primitive.E{Key: "as", Value: "game"},
		}},
	}

	pipe := mongo.Pipeline{match, group, sort, lookup}
	err := mgo.Aggregate(mgo.Acts, pipe, &acts)

	for i, v := range acts {
		acts[i].Hour = v.Act.Duration / 60
		acts[i].Min = v.Act.Duration % 60
	}
	return acts, err
}

func (s *service) ByGame(gameId string) ([]act.Act, error) {
	var acts []act.Act
	err := s.Base.FindMany(mgo.Acts, &acts, bson.D{{"game_id", format.ToObjID(gameId)}}, bson.D{})
	return acts, err
}

func (s *service) ByMonth(date string) ([]act.Detail, error) {
	if len(date) == 0 {
		return nil, fmt.Errorf("unknown date")
	}

	var month string
	if len(date) == 0 {
		month = datetime.Today(datetime.DEFAULT)[0:6]
	} else if len(date) > 6 {
		month = date[0:6]
	}

	match := bson.D{{"$match", bson.D{{"date", primitive.Regex{Pattern: "^" + month, Options: "m"}}}}}
	group := bson.D{
		{"$group", bson.D{
			{"_id", "$game_id"},
			{"type", bson.D{{"$first", "$type"}}},
			{"game_id", bson.D{{"$first", "$game_id"}}},
			{"duration", bson.D{{"$sum", "$duration"}}},
		}},
	}

	sort := bson.D{{"$sort", bson.D{{"type", 1}, {"duration", 1}}}}
	lookup := bson.D{
		{"$lookup", bson.D{
			{"from", "game"},
			{"localField", "game_id"},
			{"foreignField", "_id"},
			{"as", "game"},
		}},
	}

	var acts []act.Detail
	pine := []bson.D{match, group, sort, lookup}
	err := mgo.Aggregate(mgo.Acts, pine, &acts)
	if err != nil {
		return nil, err
	}

	for i, v := range acts {
		acts[i].Hour = v.Act.Duration / 60
		acts[i].Min = v.Act.Duration % 60
	}
	return acts, nil
}

func (s *service) Delete(id string) error {
	return s.Base.Delete(mgo.Acts, id)
}

func (s *service) Create(a act.Act) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	if a.Type == act.GAMING {
		gSrv := game.NewService()
		g := gSrv.ByID(a.GameID.Hex())
		g.PlayTime += a.Duration
		gSrv.Update(g)
	}
	return s.Base.Create(mgo.Acts, a)
}

func (s *service) DaySum(ymd string) act.Summary {
	result := []bson.M{}
	var sum act.Summary

	// Check parameter
	if len(ymd) < 8 {
		return sum
	}

	gMatch := bson.D{primitive.E{Key: "$match", Value: bson.D{
		primitive.E{Key: "date", Value: ymd},
		primitive.E{Key: "type", Value: act.GAMING}},
	}}

	pMatch := bson.D{primitive.E{Key: "$match", Value: bson.D{
		primitive.E{Key: "date", Value: ymd},
		primitive.E{Key: "type", Value: act.PROGRAMMING}},
	}}

	group := bson.D{primitive.E{Key: "$group", Value: bson.D{
		primitive.E{Key: "_id", Value: "$type"},
		primitive.E{Key: "duration", Value: bson.D{
			primitive.E{Key: "$sum", Value: "$duration"}},
		}},
	}}

	gPipe := []bson.D{gMatch, group}
	mgo.Aggregate(mgo.Acts, gPipe, &result)
	if len(result) != 0 {
		sum.GameDur = int(result[0]["duration"].(int32))
		sum.GameHour = sum.GameDur / 60
		sum.GameMin = sum.GameDur % 60
	}

	pPipe := []bson.D{pMatch, group}
	mgo.Aggregate(mgo.Acts, pPipe, &result)
	if len(result) != 0 {
		sum.PgmDur = int(result[0]["duration"].(int32))
		sum.PgmHour = sum.PgmDur / 60
		sum.PgmMin = sum.PgmDur % 60
	}

	return sum
}

func (s *service) Duration(date string, typ act.Type) int {
	if len(date) == 0 {
		return 0
	}

	result := []bson.M{}
	var match bson.D

	if len(date) == 6 {
		match = bson.D{
			primitive.E{Key: "$match", Value: bson.D{
				primitive.E{Key: "date", Value: primitive.Regex{Pattern: "^" + date, Options: "m"}},
				primitive.E{Key: "type", Value: typ},
			}},
		}
	} else {
		match = bson.D{primitive.E{Key: "$match", Value: bson.D{
			primitive.E{Key: "date", Value: date},
			primitive.E{Key: "type", Value: typ},
		}}}
	}

	group := bson.D{
		primitive.E{Key: "$group", Value: bson.D{
			primitive.E{Key: "_id", Value: "$type"},
			primitive.E{Key: "duration", Value: bson.D{
				primitive.E{Key: "$sum", Value: "$duration"}}},
		}},
	}

	pipeline := []bson.D{match, group}
	if mgo.Aggregate(mgo.Acts, pipeline, &result) != nil {
		return 0
	}

	if len(result) == 0 {
		return 0
	}

	return int(result[0]["duration"].(int32))
}

func (s *service) MonthSum(yyyymm string) act.Summary {
	result := []bson.M{}
	var sum act.Summary

	// Check parameter
	if len(yyyymm) < 6 {
		return sum
	}

	gMatch := bson.D{
		primitive.E{Key: "$match", Value: bson.D{
			primitive.E{Key: "date", Value: primitive.Regex{Pattern: "^" + yyyymm, Options: "m"}},
			primitive.E{Key: "type", Value: act.GAMING},
		}},
	}

	pMatch := bson.D{
		primitive.E{Key: "$match", Value: bson.D{
			primitive.E{Key: "date", Value: primitive.Regex{Pattern: "^" + yyyymm, Options: "m"}},
			primitive.E{Key: "type", Value: act.PROGRAMMING},
		}},
	}

	group := bson.D{
		primitive.E{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$type"},
			{Key: "duration", Value: bson.D{
				primitive.E{Key: "$sum", Value: "$duration"}}},
		}},
	}

	gPipe := []bson.D{gMatch, group}
	mgo.Aggregate(mgo.Acts, gPipe, &result)
	if len(result) != 0 {
		sum.GameDur = int(result[0]["duration"].(int32))
		sum.GameHour = sum.GameDur / 60
		sum.GameMin = sum.GameDur % 60
	}

	pPipe := []bson.D{pMatch, group}
	mgo.Aggregate(mgo.Acts, pPipe, &result)
	if len(result) != 0 {
		sum.PgmDur = int(result[0]["duration"].(int32))
		sum.PgmHour = sum.PgmDur / 60
		sum.PgmMin = sum.PgmDur % 60
	}

	return sum
}
