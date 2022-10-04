package base

import "gopkg.in/mgo.v2"

type baseService struct {
	Repo BaseRepo
}

type BaseService interface {
	Create(col *mgo.Collection, T interface{}) bool
	Delete(col *mgo.Collection, id string) error
}

func NewBaseService() BaseService {
	return &baseService{Repo: NewBaseRepo()}
}

func (srv *baseService) Create(col *mgo.Collection, T interface{}) bool {
	return srv.Repo.Create(col, T)
}

func (srv *baseService) Delete(col *mgo.Collection, id string) error {
	return srv.Repo.Delete(col, id)
}
