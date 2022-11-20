package word

import (
	"ditto/db/mgo"
	"ditto/model/word"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type wordService struct{}

type WordService interface {
	ByID(id string) word.Word
	PageIsChecked(isChecked int, page, limit int) ([]word.Word, int)
	Check(id string, isChecked int) error
	Create(word word.Word) error
	Update(word word.Word) error
	Delete(id string) error
}

func NewWordService() WordService {
	return &wordService{}
}

func (s *wordService) ByID(id string) word.Word {
	var w word.Word
	result := mgo.FindID(mgo.Words, id)
	result.Decode(&w)
	return w
}

func (s *wordService) PageIsChecked(isChecked int, page, limit int) ([]word.Word, int) {
	var filter bson.D

	var words []word.Word
	sort := bson.D{primitive.E{Key: "created_at", Value: -1}}

	totalPages, err := mgo.FindPage(mgo.Words, &words, filter, page, limit, sort)
	if err != nil {
		return nil, 1
	}

	return words, totalPages
}

func (s *wordService) Check(id string, isChecked int) error {
	update := bson.D{primitive.E{
		Key: "$set", Value: bson.D{
			primitive.E{Key: "is_checked", Value: isChecked},
			primitive.E{Key: "updated_at", Value: time.Now()},
		}},
	}
	return mgo.Update(mgo.Words, id, update)
}

func (s *wordService) Create(w word.Word) error {
	w.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Words, w)
}

func (s *wordService) Delete(id string) error {
	err := mgo.DeleteID(mgo.Words, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *wordService) Update(w word.Word) error {
	update := bson.D{primitive.E{
		Key: "$set", Value: bson.D{
			primitive.E{Key: "word", Value: w.Word},
			primitive.E{Key: "example", Value: w.Example},
			primitive.E{Key: "meaning", Value: w.Meaning},
			primitive.E{Key: "is_checked", Value: w.IsChecked},
			primitive.E{Key: "updated_at", Value: time.Now()},
		}},
	}
	return mgo.Update(mgo.Words, w.ID, update)
}
