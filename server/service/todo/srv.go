package todo

import (
	"ditto/model/todo"
	"ditto/service/base"
)

var col = "todo"

type service struct {
	Base base.BaseRepo
	Repo Repo
}

type Service interface {
	All() []todo.Todo
	ByID(id string) todo.Todo
	Check(id string, isChecked int)
	Create(todo todo.Todo)
	DelChecked()
}

func NewService() Service {
	return &service{
		Repo: NewRepo(),
	}
}

func (s *service) All() []todo.Todo {
	return s.Repo.All()
}

func (s *service) ByID(id string) todo.Todo {
	return s.Repo.ByID(id)
}

func (s *service) Check(id string, isChecked int) {
	t := s.ByID(id)
	if isChecked == 1 {
		t.IsChecked = true
	} else {
		t.IsChecked = false
	}
	s.Repo.Check(t)
}

func (s *service) Create(todo todo.Todo) {
	s.Repo.Create(todo)
}

func (s *service) DelChecked() {
	s.Repo.DelChecked()
}
