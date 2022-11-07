package user

import (
	"ditto/db/mgo"
	"ditto/model/user"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type UserService interface {
	ByUsername(username string) user.User
	Login(username, password string) (user.User, error)
	Register(u user.User) error
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) ByUsername(username string) user.User {
	var u user.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	mgo.FindOne(mgo.Users, filter).Decode(&u)
	return u
}

func (s *userService) Login(username, password string) (user.User, error) {
	u := s.ByUsername(username)
	if len(u.Username) > 0 && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
		return u, nil
	} else {
		return user.User{}, fmt.Errorf("login failed")
	}
}

func (s *userService) Register(u user.User) error {
	find := s.ByUsername(u.Username)
	if len(find.Username) == 0 {
		return mgo.Insert(mgo.Users, u)
	} else {
		return fmt.Errorf("register failed")
	}
}
