package word

import (
	"ditto/db/mgo"
	"ditto/model/word"
	"ditto/service/base"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type wordService struct {
	Base base.BaseRepo
	Repo WordRepo
}

type WordService interface {
	ByID(id string) word.Word
	ByDate(date string, isCheck int) []word.Word
	ByIsChecked(isCheck int) []word.Word
	Check(isCheck int, id string) error
	Create(word word.Word) error
	Update(word word.Word) error
	Delete(id string) error
}

func NewWordService() WordService {
	return &wordService{
		Base: base.NewBaseRepo(),
		Repo: NewWordRepo(),
	}
}

func (s *wordService) ByDate(date string, isCheck int) []word.Word {
	return s.Repo.ByDate(date, isCheck)
}

func (s *wordService) ByID(id string) word.Word {
	return s.Repo.ByID(id)
}

func (s *wordService) ByIsChecked(isCheck int) []word.Word {
	return s.Repo.ByIsChecked(isCheck)
}

func (s *wordService) Check(isCheck int, id string) error {
	return s.Repo.Check(isCheck, id)
}

func (s *wordService) Create(w word.Word) error {
	w.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Words, w)
}

func (s *wordService) Delete(id string) error {
	return s.Base.Delete(mgo.Words, id)
}

func (s *wordService) Update(w word.Word) error {
	return s.Repo.Update(w)
}
