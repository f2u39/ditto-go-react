package game

import (
	"ditto/db/mongo"
	"ditto/model/game"
	"ditto/service/inc"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type repo struct{}

type Repo interface {
	byGenre(genre game.Genre) []game.Game
	byID(id string) game.Game
	byStatus(status game.Status) []game.Detail
	count(status game.Status) int
	counts() (int, int, int)
	create(g game.Game) bool
	delete(id string) error
	pageByStatus(status game.Status, platform game.Platform, page, limit int) ([]game.Detail, int)
	update(g game.Game) error
}

func NewRepo() Repo {
	return &repo{}
}

func (*repo) byID(id string) game.Game {
	g := game.Game{}
	mongo.FindByID(mongo.Games, id, &g)
	return g
}

func (*repo) byGenre(genre game.Genre) []game.Game {
	var games []game.Game
	mongo.FindMany(mongo.Games, &games, bson.M{"genre": genre}, "title")
	return games
}

func (*repo) byStatus(status game.Status) []game.Detail {
	var qry bson.M

	// Default status is "playing"
	if len(status) != 0 {
		qry = bson.M{"status": status}
	} else {
		qry = bson.M{"status": game.PLAYING}
	}

	var games []game.Game
	mongo.FindMany(mongo.Games, &games, qry, "title")

	incSrv := inc.NewIncService()
	var details []game.Detail
	for _, g := range games {
		detail := game.Detail{
			Game:      g,
			Developer: incSrv.ByID(g.DeveloperID.Hex()),
			Publisher: incSrv.ByID(g.PublisherID.Hex()),
			PlayHour:  g.PlayTime / 60,
			PlayMin:   g.PlayTime % 60,
		}
		details = append(details, detail)
	}
	return details
}

func (*repo) count(status game.Status) int {
	qry := bson.M{"status": status}
	cnt, _ := mongo.Count(mongo.Games, qry)
	return cnt
}

func (r *repo) counts() (int, int, int) {
	playedCnt := r.count(game.PLAYED)
	playingCnt := r.count(game.PLAYING)
	blockingCnt := r.count(game.BLOCKING)
	return playedCnt, playingCnt, blockingCnt
}

func (*repo) create(g game.Game) bool {
	g.ID = bson.NewObjectId()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return mongo.Insert(mongo.Games, g)
}

func (*repo) delete(id string) error {
	err := mongo.DeleteByID(mongo.Games, id)
	if err != nil {
		return err
	}
	return nil
}

func (*repo) pageByStatus(status game.Status, platform game.Platform, page, limit int) ([]game.Detail, int) {
	var qry bson.M

	// Check status
	if len(status) != 0 {
		qry = bson.M{"status": status}
	} else {
		// If empty, the default status is "playing"
		qry = bson.M{"status": game.PLAYING}
	}

	if len(platform) != 0 && platform != "All" {
		qry["platform"] = platform
	}

	var games []game.Game
	totalPages, err := mongo.FindPage(mongo.Games, &games, qry, page, limit, "title")
	if err != nil {
		return nil, 1
	}

	incSrv := inc.NewIncService()
	var details []game.Detail
	for _, g := range games {
		detail := game.Detail{
			Game:      g,
			Developer: incSrv.ByID(g.DeveloperID.Hex()),
			Publisher: incSrv.ByID(g.PublisherID.Hex()),
			PlayHour:  g.PlayTime / 60,
			PlayMin:   g.PlayTime % 60,
		}
		details = append(details, detail)
	}
	return details, totalPages
}

func (*repo) update(g game.Game) error {
	return mongo.Update(mongo.Games, g.ID, g)
}
