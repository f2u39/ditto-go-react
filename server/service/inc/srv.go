package inc

import (
	"ditto/db/mgo"
	"ditto/model/inc"
	"ditto/service/base"
)

type IncService interface {
	All() []inc.Inc
	ByID(id string) inc.Inc
	Create(inc inc.Inc) bool
	Developers() []inc.Inc
	Publishers() []inc.Inc
	Update(inc inc.Inc) error
}

func NewIncService() IncService {
	return &incService{
		Base: base.NewBaseRepo(),
		Repo: NewIncRepo()}
}

type incService struct {
	Base base.BaseRepo
	Repo IncRepo
}

func (s *incService) All() []inc.Inc {
	return s.Repo.All()
}

func (s *incService) ByID(id string) inc.Inc {
	return s.Repo.ByID(id)
}

func (s *incService) Developers() []inc.Inc {
	return s.Repo.Developers()
}

func (s *incService) Publishers() []inc.Inc {
	return s.Repo.Publishers()
}

func (s *incService) Create(inc inc.Inc) bool {
	mgo.Before(&inc.ID, &inc.CreatedAt, &inc.UpdatedAt)
	return s.Base.Create(mgo.Incs, inc)
}

func (s *incService) Update(inc inc.Inc) error {
	return s.Base.Update(mgo.Incs, inc.ID, inc)
}
