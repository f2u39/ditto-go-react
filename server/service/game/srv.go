package game

import (
	"ditto/db/mgo"
	"ditto/model/game"
	"ditto/service/base"
	"ditto/service/inc"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameService interface {
	ByGenre(genre game.Genre) []game.Game
	ByID(id string) game.Game
	ByPlaying() []game.Game
	ByStatus(status game.Status) []game.Detail
	Counts() (int, int, int)
	Create(g game.Game) error
	Delete(id string) error
	PageByStatus(status game.Status, platform game.Platform, page, limit int) ([]game.Detail, int)
	Update(g game.Game) error
}

func NewService() GameService {
	return &service{Base: base.NewBaseRepo()}
}

type service struct {
	Base base.BaseRepo
}

func (s *service) ByID(id string) game.Game {
	var game game.Game
	result, err := mgo.FindID(mgo.Games, id)
	if err != nil {
		return game
	}
	result.Decode(&game)
	return game
}

func (s *service) ByGenre(genre game.Genre) []game.Game {
	var games []game.Game
	filter := bson.D{primitive.E{Key: "genre", Value: genre}}
	sort := bson.D{primitive.E{Key: "title", Value: 1}}
	mgo.FindMany(mgo.Games, &games, filter, sort)
	return games
}

func (s *service) ByPlaying() []game.Game {
	var games []game.Game
	filter := bson.D{primitive.E{Key: "status", Value: game.PLAYING}}
	sort := bson.D{primitive.E{Key: "title", Value: 1}}
	mgo.FindMany(mgo.Games, &games, filter, sort)
	return games
}

func (s *service) ByStatus(status game.Status) []game.Detail {
	var filter bson.D
	var games []game.Game

	// Default status is "playing"
	if len(status) != 0 {
		filter = bson.D{primitive.E{Key: "status", Value: status}}
	} else {
		filter = bson.D{primitive.E{Key: "status", Value: game.PLAYING}}
	}

	sort := bson.D{primitive.E{Key: "title", Value: 1}}
	mgo.FindMany(mgo.Games, &games, filter, sort)

	incSrv := inc.NewIncService()
	var details []game.Detail
	for _, g := range games {
		detail := game.Detail{
			Game:      g,
			Developer: incSrv.ByID(g.DeveloperID),
			Publisher: incSrv.ByID(g.PublisherID),
			PlayHour:  g.PlayTime / 60,
			PlayMin:   g.PlayTime % 60,
		}
		details = append(details, detail)
	}
	return details
}

func count(status game.Status) int {
	filter := bson.D{primitive.E{Key: "status", Value: status}}
	cnt, _ := mgo.Count(mgo.Games, filter)
	return int(cnt)
}

func (s *service) Counts() (int, int, int) {
	playedCnt := count(game.PLAYED)
	playingCnt := count(game.PLAYING)
	blockingCnt := count(game.BLOCKING)
	return playedCnt, playingCnt, blockingCnt
}

func (s *service) Create(g game.Game) error {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Games, g)
}

func (s *service) Delete(id string) error {
	err := mgo.DeleteID(mgo.Games, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) PageByStatus(status game.Status, platform game.Platform, page, limit int) ([]game.Detail, int) {
	var filter bson.D

	// Check status
	if len(status) != 0 {
		filter = bson.D{primitive.E{Key: "status", Value: status}}
	} else {
		// If empty, the default status is "playing"
		filter = bson.D{primitive.E{Key: "status", Value: game.PLAYING}}
	}

	if len(platform) != 0 && platform != "All" {
		filter = append(filter, primitive.E{Key: "platform", Value: platform})
	}

	var games []game.Game
	sort := bson.D{primitive.E{Key: "title", Value: 1}}
	totalPages, err := mgo.FindPage(mgo.Games, &games, filter, page, limit, sort)
	if err != nil {
		return nil, 1
	}

	incSrv := inc.NewIncService()
	var details []game.Detail

	for _, g := range games {
		detail := game.Detail{
			Game:      g,
			Developer: incSrv.ByID(g.DeveloperID),
			Publisher: incSrv.ByID(g.PublisherID),
			PlayHour:  g.PlayTime / 60,
			PlayMin:   g.PlayTime % 60,
		}
		details = append(details, detail)
	}
	return details, totalPages
}

func (s *service) Update(g game.Game) error {
	return mgo.Update(mgo.Games, g.ID, g)
}
