package user

import "ditto/model/user"

type userService struct {
	Repo UserRepo
}

type UserService interface {
	ByUsername(username string) (user.User, bool)
	Login(username, password string) (user.User, bool)
	Register(u user.User) (user.User, bool)
}

func NewUserService() UserService {
	return &userService{Repo: NewUserRepo()}
}

func (s *userService) ByUsername(username string) (user.User, bool) {
	return s.Repo.ByUsername(username)
}

func (s *userService) Login(username, password string) (user.User, bool) {
	return s.Repo.Login(username, password)
}

func (s *userService) Register(u user.User) (user.User, bool) {
	return s.Repo.Register(u)
}
