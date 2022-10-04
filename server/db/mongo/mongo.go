package mongo

import (
	"ditto/lib/err"
	"ditto/lib/format"
	"ditto/model/config"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	Sess *mgo.Session

	Acts    *mgo.Collection
	Animes  *mgo.Collection
	Incs    *mgo.Collection
	Games   *mgo.Collection
	Users   *mgo.Collection
	Todos   *mgo.Collection
	Words   *mgo.Collection
	Studios *mgo.Collection
)

func connect() (*mgo.Session, error) {
	// Connect to MongoDB
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{config.Config.MongoDB.URL},
		Database: config.Config.MongoDB.Database,
		// Username: dbCfg.MongoDB.Username,
		// Password: dbCfg.MongoDB.Password,
	}
	return mgo.DialWithInfo(dialInfo)
}

func Before(id *bson.ObjectId, createdAt, updatedAt *time.Time) {
	*id = bson.NewObjectId()
	*createdAt = time.Now()
	*updatedAt = time.Now()
}

func Init() {
	var err error
	Sess, err = connect()
	if err != nil {
		panic(err)
	}

	// In the Monotonic consistency mode reads may not be entirely up-to-date,
	// but they will always see the history of changes moving forward,
	// the data read will be consistent across sequential queries in the same session,
	// and modifications made within the session will be observed in following queries (read-your-writes).
	Sess.SetMode(mgo.Monotonic, true)

	db := config.Config.MongoDB.Database

	Acts = Sess.DB(db).C("act")
	Animes = Sess.DB(db).C("anime")
	Incs = Sess.DB(db).C("inc")
	Games = Sess.DB(db).C("game")
	Users = Sess.DB(db).C("user")
	Todos = Sess.DB(db).C("todo")
	Words = Sess.DB(db).C("word")
	Studios = Sess.DB(db).C("studio")
}

func close() {
	Sess.Close()
}

func Count(col *mgo.Collection, qry bson.M) (int, error) {
	return col.Find(qry).Count()
}

func Delete(col *mgo.Collection, T interface{}) error {
	return col.Remove(T)
}

func DeleteById(col *mgo.Collection, id interface{}) error {
	switch id.(type) {
	case string:
		oid := format.ToObjId(fmt.Sprintf("%v", id))
		return col.RemoveId(oid)
	case bson.ObjectId:
		return col.RemoveId(id)
	default:
		return nil
	}
}

func Insert(col *mgo.Collection, T interface{}) bool {
	return err.E(col.Insert(T))
}

func FindOne(col *mgo.Collection, qry bson.M, T interface{}) error {
	return col.Find(qry).One(T)
}

func FindID(col *mgo.Collection, id interface{}, T interface{}) error {
	switch id.(type) {
	case string:
		oid := format.ToObjId(fmt.Sprintf("%v", id))
		return col.FindId(oid).One(T)
	case bson.ObjectId:
		return col.FindId(id).One(T)
	default:
		return fmt.Errorf("id is not string or bson.ObjectId")
	}
}

func LookUp(col *mgo.Collection, pipeline []bson.M, T interface{}) error {
	return col.Pipe(pipeline).All(T)
}

func FindMany(col *mgo.Collection, T interface{}, qry bson.M, sorts ...string) error {
	if len(sorts) == 0 {
		return col.Find(qry).All(T)
	} else {
		return col.Find(qry).Sort(sorts...).All(T)
	}
}

// Return total pages and error
func FindPage(col *mgo.Collection, T any, qry bson.M, page, limit int, sorts ...string) (int, error) {
	// Check is page or limit is 0
	if page <= 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	cnt, err := col.Find(qry).Count()
	if err != nil {
		return 0, err
	}

	totalPages := cnt / limit
	if cnt%limit > 0 {
		totalPages += 1
	}
	if totalPages == 0 {
		totalPages = 1
	}

	if len(sorts) == 0 {
		return totalPages, col.Find(qry).Skip((page - 1) * limit).Limit(limit).All(T)
	} else {
		return totalPages, col.Find(qry).Sort(sorts...).Skip((page - 1) * limit).Limit(limit).All(T)
	}
}

func Update(col *mgo.Collection, id interface{}, T interface{}) error {
	return col.UpdateId(id, T)
}
