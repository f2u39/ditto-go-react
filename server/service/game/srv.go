package game

import (
	"ditto/db/mgo"
	"ditto/model/game"
	"ditto/service/base"

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
	return &service{
		Base: base.NewBaseRepo(),
		Repo: NewRepo()}
}

type service struct {
	Base base.BaseRepo
	Repo Repo
}

func (s *service) ByID(id string) game.Game {
	return s.Repo.byID(id)
}

func (s *service) ByGenre(genre game.Genre) []game.Game {
	return s.Repo.byGenre(genre)
}

func (s *service) ByPlaying() []game.Game {
	var games []game.Game
	s.Base.FindMany(mgo.Games, &games, bson.D{primitive.E{Key: "status", Value: game.PLAYING}}, bson.D{primitive.E{Key: "title", Value: 1}})
	return games
}

func (s *service) ByStatus(status game.Status) []game.Detail {
	return s.Repo.byStatus(status)
}

func (s *service) Counts() (int, int, int) {
	return s.Repo.counts()
}

func (s *service) Create(g game.Game) error {
	return s.Repo.create(g)
}

func (s *service) Delete(id string) error {
	return s.Repo.delete(id)
}

func (s *service) PageByStatus(status game.Status, platform game.Platform, page, limit int) ([]game.Detail, int) {
	return s.Repo.pageByStatus(status, platform, page, limit)
}

func (s *service) Update(g game.Game) error {
	return s.Repo.update(g)
}
