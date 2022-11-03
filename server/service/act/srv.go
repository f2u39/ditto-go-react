package act

import (
	"ditto/db/mgo"
	"ditto/lib/format"
	"ditto/model/act"
	"ditto/service/base"
	"ditto/service/game"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	Repo Repo
}

func NewService() Service {
	return &service{
		Base: base.NewBaseRepo(),
		Repo: NewRepo()}
}

func (s *service) ByDate(date string) ([]act.Detail, error) {
	return s.Repo.ByDate(date)
}

func (s *service) ByGame(gameId string) ([]act.Act, error) {
	var acts []act.Act
	err := s.Base.FindMany(mgo.Acts, &acts, bson.D{{"game_id", format.ObjId(gameId)}}, "")
	return acts, err
}

func (s *service) ByMonth(date string) ([]act.Detail, error) {
	return s.Repo.ByMonth(date)
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
	return s.Repo.DaySum(ymd)
}

func (s *service) Duration(ymd string, typ act.Type) int {
	return s.Repo.Duration(ymd, typ)
}

func (s *service) MonthSum(yyyymm string) act.Summary {
	return s.Repo.MonthSum(yyyymm)
}
