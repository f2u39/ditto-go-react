package base

import "go.mongodb.org/mongo-driver/mongo"

type baseService struct {
	Repo BaseRepo
}

type BaseService interface {
	Delete(col *mongo.Collection, id any) error
}

func NewBaseService() BaseService {
	return &baseService{Repo: NewBaseRepo()}
}

func (srv *baseService) Delete(col *mongo.Collection, id any) error {
	return srv.Repo.Delete(col, id)
}
