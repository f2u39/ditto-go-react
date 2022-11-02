package word

import (
	"ditto/db/mgo"
	"ditto/model/word"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type wordRepo struct{}

type WordRepo interface {
	ByDate(date string, isCheck int) []word.Word
	ByID(id string) word.Word
	ByIsChecked(isCheck int) []word.Word
	Check(isCheck int, id string) error
	Create(word word.Word) bool
	Update(word word.Word) error
}

func NewWordRepo() WordRepo {
	return &wordRepo{}
}

func (*wordRepo) ByDate(date string, isCheck int) []word.Word {
	var words []word.Word
	match := bson.M{"$match": bson.M{"date": date, "is_checked": isCheck}}
	sort := bson.M{"$sort": bson.M{"word": 1}}
	pipeline := []bson.M{match, sort}
	mgo.LookUp(mgo.Words, pipeline, &words)
	return words
}

func (*wordRepo) ByIsChecked(isCheck int) []word.Word {
	var words []word.Word
	match := bson.M{"$match": bson.M{"is_checked": isCheck}}
	sort := bson.M{"$sort": bson.M{"created_at": -1}}
	pipeline := []bson.M{match, sort}
	mgo.LookUp(mgo.Words, pipeline, &words)
	return words
}

func (*wordRepo) ByID(id string) word.Word {
	w := word.Word{}
	mgo.FindID(mgo.Words, id, &w)
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

func (*wordRepo) Create(word word.Word) bool {
	word.ID = bson.NewObjectId()
	word.CreatedAt = time.Now()
	word.UpdatedAt = time.Now()
	return mgo.Insert(mgo.Words, word)
}

func (*wordRepo) Update(word word.Word) error {
	return mgo.Update(mgo.Words, word.ID, word)
}
