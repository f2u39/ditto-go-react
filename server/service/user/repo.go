package user

import (
	"ditto/db/mgo"
	"ditto/model/user"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type userRepo struct{}

type UserRepo interface {
	ByUsername(username string) (user.User, bool)
	Login(username, password string) (user.User, bool)
	Register(u user.User) (user.User, bool)
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (*userRepo) ByUsername(username string) (user.User, bool) {
	var u user.User
	qry := bson.M{"username": username}
	err := mgo.FindOne(mgo.Users, qry, &u)
	if err != nil {
		log.Println(err)
	}
	return u, err == nil
}

func (r *userRepo) Login(username, password string) (user.User, bool) {
	// var u user.User
	// qry := bson.M{"username": username, "password": password}
	// err := mongo.FindOne(mongo.Users, qry, &u)
	// if err != nil {
	// 	return user.User{}, false
	// }
	// return u, true

	u, ok := r.ByUsername(username)
	if ok && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
		return u, true
	} else {
		return user.User{}, false
	}
}

func (r *userRepo) Register(user user.User) (user.User, bool) {
	u, ok := r.ByUsername(user.Username)
	if !ok && len(u.Password) == 0 {
		return u, mgo.Insert(mgo.Users, user)
	} else {
		return user, false
	}
}
