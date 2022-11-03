package user

import "ditto/model/user"

type userService struct {
	Repo UserRepo
}

type UserService interface {
	ByUsername(username string) (user.User, error)
	Login(username, password string) (user.User, error)
	Register(u user.User) (user.User, error)
}

func NewUserService() UserService {
	return &userService{Repo: NewUserRepo()}
}

func (s *userService) ByUsername(username string) (user.User, error) {
	return s.Repo.ByUsername(username)
}

func (s *userService) Login(username, password string) (user.User, error) {
	return s.Repo.Login(username, password)
}

func (s *userService) Register(u user.User) (user.User, error) {
	return s.Repo.Register(u)
}
