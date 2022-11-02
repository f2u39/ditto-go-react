package act

import (
	"ditto/db/mgo"
	"ditto/lib/datetime"
	"ditto/model/act"
	"fmt"

	"gopkg.in/mgo.v2/bson"
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
	match := bson.D{{"$match", bson.D{{"date", date}}}}

	group := bson.D{"$group", bson.D{
		"_id", "$_id",
		"type", bson.D{"$first", "$type"},
		"game_id", bson.D{"$first", "$game_id"},
		"duration", bson.D{"$sum", "$duration"},
	}}
	sort := bson.D{"$sort", bson.D{"type": 1, "duration": 1}}
	lookup := bson.D{"$lookup", bson.D{
		"from", "game",
		"localField", "game_id",
		"foreignField", "_id",
		"as", "game",
	}}
	pipeline := []bson.D{match, group, sort, lookup}
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

	match := bson.M{"$match": bson.M{"date": bson.RegEx{Pattern: "^" + month, Options: "m"}}}
	group := bson.M{"$group": bson.M{
		"_id":      "$game_id",
		"type":     bson.M{"$first": "$type"},
		"game_id":  bson.M{"$first": "$game_id"},
		"duration": bson.M{"$sum": "$duration"},
	}}
	sort := bson.M{"$sort": bson.M{"type": 1, "duration": 1}}
	lookup := bson.M{"$lookup": bson.M{
		"from":         "game",
		"localField":   "game_id",
		"foreignField": "_id",
		"as":           "game",
	}}

	var acts []act.Detail
	pipeline := []bson.M{match, group, sort, lookup}
	err := mgo.LookUp(mgo.Acts, pipeline, &acts)
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
	result := []bson.M{}
	var sum act.Summary

	// Check parameter
	if len(ymd) < 8 {
		return sum
	}

	gMatch := bson.M{"$match": bson.M{"date": ymd, "type": act.GAMING}}
	pMatch := bson.M{"$match": bson.M{"date": ymd, "type": act.PROGRAMMING}}

	group := bson.M{"$group": bson.M{
		"_id":      "$type",
		"duration": bson.M{"$sum": "$duration"},
	}}

	gPipe := []bson.M{gMatch, group}
	mgo.LookUp(mgo.Acts, gPipe, &result)
	if len(result) != 0 {
		sum.GameDur = result[0]["duration"].(int)
		sum.GameHour = sum.GameDur / 60
		sum.GameMin = sum.GameDur % 60
	}

	pPipe := []bson.M{pMatch, group}
	mgo.LookUp(mgo.Acts, pPipe, &result)
	if len(result) != 0 {
		sum.PgmDur = result[0]["duration"].(int)
		sum.PgmHour = sum.PgmDur / 60
		sum.PgmMin = sum.PgmDur % 60
	}

	return sum
}

func (*repo) Duration(date string, typ act.Type) int {
	if len(date) == 0 {
		return 0
	}

	result := []bson.M{}
	var match bson.M
	if len(date) == 6 {
		match = bson.M{"$match": bson.M{"date": bson.RegEx{Pattern: "^" + date, Options: "m"}, "type": typ}}
	} else {
		match = bson.M{"$match": bson.M{"date": date, "type": typ}}
	}

	group := bson.M{"$group": bson.M{
		"_id":      "$type",
		"duration": bson.M{"$sum": "$duration"},
	}}

	pipeline := []bson.M{match, group}
	if mgo.LookUp(mgo.Acts, pipeline, &result) != nil {
		return 0
	}

	if len(result) == 0 {
		return 0
	}
	return result[0]["duration"].(int)
}

func (*repo) MonthSum(yyyymm string) act.Summary {
	result := []bson.M{}
	var sum act.Summary

	// Check parameter
	if len(yyyymm) < 6 {
		return sum
	}

	gMatch := bson.M{"$match": bson.M{"date": bson.RegEx{Pattern: "^" + yyyymm, Options: "m"}, "type": act.GAMING}}
	pMatch := bson.M{"$match": bson.M{"date": bson.RegEx{Pattern: "^" + yyyymm, Options: "m"}, "type": act.PROGRAMMING}}

	group := bson.M{"$group": bson.M{
		"_id":      "$type",
		"duration": bson.M{"$sum": "$duration"},
	}}

	gPipe := []bson.M{gMatch, group}
	mgo.LookUp(mgo.Acts, gPipe, &result)
	if len(result) != 0 {
		sum.GameDur = result[0]["duration"].(int)
		sum.GameHour = sum.GameDur / 60
		sum.GameMin = sum.GameDur % 60
	}

	pPipe := []bson.M{pMatch, group}
	mgo.LookUp(mgo.Acts, pPipe, &result)
	if len(result) != 0 {
		sum.PgmDur = result[0]["duration"].(int)
		sum.PgmHour = sum.PgmDur / 60
		sum.PgmMin = sum.PgmDur % 60
	}

	return sum
}
