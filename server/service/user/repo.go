package user

import (
	"ditto/db/mgo"
	"ditto/model/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct{}

type UserRepo interface {
	ByUsername(username string) (user.User, error)
	Login(username, password string) (user.User, error)
	Register(u user.User) (user.User, error)
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (*userRepo) ByUsername(username string) (user.User, error) {
	var u user.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	err := mgo.FindOne(mgo.Users, filter, &u)
	return u, err
}

func (r *userRepo) Login(username, password string) (user.User, error) {
	u, err := r.ByUsername(username)
	if err != nil && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
		return u, err
	} else {
		return user.User{}, err
	}
}

func (r *userRepo) Register(user user.User) (user.User, error) {
	u, err := r.ByUsername(user.Username)
	if err != nil && len(u.Password) == 0 {
		return u, mgo.Insert(mgo.Users, user)
	} else {
		return user, nil
	}
}
