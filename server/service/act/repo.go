package act

import (
	"ditto/db/mgo"
	"ditto/lib/datetime"
	"ditto/model/act"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct{}

type Repo interface {
	ByDate(date string) ([]act.Detail, error)
	// Create(act act.Act) bool
	ByMonth(date string) ([]act.Detail, error)
	DaySum(ymd string) act.Summary
	Duration(ymd string, typ act.Type) int
	MonthSum(yyyymm string) act.Summary
}

func NewRepo() Repo {
	return &repo{}
}

func (*repo) ByDate(date string) ([]act.Detail, error) {
	if len(date) == 0 {
		return nil, fmt.Errorf("unknown date")
	}

	var acts []act.Detail
	match := bson.D{
		{"$match", bson.D{{"date", date}}},
	}

	group := bson.D{
		{"$group", bson.D{
			{"_id", "$_id"},
			{"type", bson.D{{"$first", "$type"}}},
			{"game_id", bson.D{{"$first", "$game_id"}}},
			{"duration", bson.D{{"$sum", "$duration"}}},
		}},
	}

	sort := bson.D{
		{"$sort", bson.D{
			{"type", 1},
			{"duration", 1},
		}},
	}

	lookup := bson.D{
		{"$lookup", bson.D{
			{"from", "game"},
			{"localField", "game_id"},
			{"foreignField", "_id"},
			{"as", "game"},
		}},
	}

	pipeline := mongo.Pipeline{match, group, sort, lookup}
	err := mgo.Aggregate(mgo.Acts, pipeline, &acts)

	for i, v := range acts {
		acts[i].Hour = v.Act.Duration / 60
		acts[i].Min = v.Act.Duration % 60
	}
	return acts, err
}

func (*repo) ByMonth(date string) ([]act.Detail, error) {
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

func (*repo) DaySum(ymd string) act.Summary {
	result := []bson.D{}
	var sum act.Summary

	// Check parameter
	if len(ymd) < 8 {
		return sum
	}

	gMatch := bson.D{{"$match", bson.D{{"date", ymd}, {"type", act.GAMING}}}}
	pMatch := bson.D{{"$match", bson.D{{"date", ymd}, {"type", act.PROGRAMMING}}}}
	group := bson.D{
		{"$group", bson.D{
			{"_id", "$type"},
			{"duration", bson.D{{"$sum", "$duration"}}},
		}},
	}

	gPipe := []bson.D{gMatch, group}
	mgo.Aggregate(mgo.Acts, gPipe, &result)
	if len(result) != 0 {
		// TODO invalid operation: result[0]["duration"] (variable of type primitive.E) is not an interface
		// sum.GameDur = result[0]["duration"].(int)
		sum.GameHour = sum.GameDur / 60
		sum.GameMin = sum.GameDur % 60
	}

	pPipe := []bson.D{pMatch, group}
	mgo.Aggregate(mgo.Acts, pPipe, &result)
	if len(result) != 0 {
		// TODO invalid operation: result[0]["duration"] (variable of type primitive.E) is not an interface
		// sum.PgmDur = result[0]["duration"].(int)
		sum.PgmHour = sum.PgmDur / 60
		sum.PgmMin = sum.PgmDur % 60
	}

	return sum
}

func (*repo) Duration(date string, typ act.Type) int {
	if len(date) == 0 {
		return 0
	}

	result := []bson.D{}
	var match bson.D

	if len(date) == 6 {
		// TODO Need to fix
		// match = bson.D{
		// 	{"$match", bson.D{
		// 		{"date", primitive.Regex{Pattern: "^" + date, Options: "m"}},
		// 		{"type", typ},
		// 	}},
		// }
	} else {
		match = bson.D{
			{"$match", bson.D{
				{"date", date},
				{"type", typ},
			}},
		}
	}

	group := bson.D{
		{"$group", bson.D{
			{"_id", "$type"},
			{"duration", bson.D{{"$sum", "$duration"}}},
		}},
	}

	pipeline := []bson.D{match, group}
	if mgo.Aggregate(mgo.Acts, pipeline, &result) != nil {
		return 0
	}

	if len(result) == 0 {
		return 0
	}

	// TODO Need to fix
	// return result[0]["duration"].(int)
	return 0
}

func (*repo) MonthSum(yyyymm string) act.Summary {
	result := []bson.D{}
	var sum act.Summary

	// Check parameter
	if len(yyyymm) < 6 {
		return sum
	}

	gMatch := bson.D{
		{"$match", bson.D{
			{"date", primitive.Regex{Pattern: "^" + yyyymm, Options: "m"}},
			{"type", act.GAMING},
		}},
	}

	pMatch := bson.D{
		{"$match", bson.D{
			{"date", primitive.Regex{Pattern: "^" + yyyymm, Options: "m"}},
			{"type", act.PROGRAMMING},
		}},
	}

	group := bson.D{
		{"$group", bson.D{
			{"_id", "$type"},
			{"duration", bson.D{{"$sum", "$duration"}}},
		}},
	}

	gPipe := []bson.D{gMatch, group}
	mgo.Aggregate(mgo.Acts, gPipe, &result)
	if len(result) != 0 {
		// TODO Need to fix
		// sum.GameDur = result[0]["duration"].(int)
		sum.GameHour = sum.GameDur / 60
		sum.GameMin = sum.GameDur % 60
	}

	pPipe := []bson.D{pMatch, group}
	mgo.Aggregate(mgo.Acts, pPipe, &result)
	if len(result) != 0 {
		// TODO Need to fix
		// sum.PgmDur = result[0]["duration"].(int)
		sum.PgmHour = sum.PgmDur / 60
		sum.PgmMin = sum.PgmDur % 60
	}

	return sum
}
