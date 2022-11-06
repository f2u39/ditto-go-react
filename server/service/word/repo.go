package word

import (
	"ditto/db/mgo"
	"ditto/model/word"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type wordRepo struct{}

type WordRepo interface {
	ByDate(date string, isCheck int) []word.Word
	ByID(id string) word.Word
	ByIsChecked(isCheck int) []word.Word
	Check(isCheck int, id string) error
	Create(word word.Word) error
	Update(word word.Word) error
}

func NewWordRepo() WordRepo {
	return &wordRepo{}
}

func (*wordRepo) ByDate(date string, isCheck int) []word.Word {
	var words []word.Word
	match := bson.D{{"$match", bson.D{{"date", date}, {"is_checked", isCheck}}}}
	sort := bson.D{{"$sort", bson.D{{"word", 1}}}}
	pipe := []bson.D{match, sort}
	mgo.Aggregate(mgo.Words, pipe, &words)
	return words
}

func (*wordRepo) ByIsChecked(isCheck int) []word.Word {
	var words []word.Word
	match := bson.D{{"$match", bson.D{{"is_checked", isCheck}}}}
	sort := bson.D{{"$sort", bson.D{{"created_at", -1}}}}
	pipe := []bson.D{match, sort}
	mgo.Aggregate(mgo.Words, pipe, &words)
	return words
}

func (*wordRepo) ByID(id string) word.Word {
	var w word.Word
	result, err := mgo.FindID(mgo.Words, id)
	if err != nil {
		return w
	}
	result.Decode(&w)
	return w
}

func (r *wordRepo) Check(isDone int, id string) error {
	if isDone != 0 && isDone != 1 {
		return fmt.Errorf("illegal parameter isDone")
	}

	w := r.ByID(id)
	// if err != nil {
	// 	return false
	// }

	w.IsChecked = isDone
	return mgo.Update(mgo.Words, w.ID, w)
}

func (*wordRepo) Create(word word.Word) error {
	word.CreatedAt = time.Now()
	word.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Words, word)
}

func (*wordRepo) Update(word word.Word) error {
	return mgo.Update(mgo.Words, word.ID, word)
}
