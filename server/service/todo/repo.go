package todo

import (
	"ditto/db/firebase"
	"ditto/lib/datetime"
	"ditto/model/todo"
	"log"
)

type repo struct{}

type Repo interface {
	All() []todo.Todo
	ByID(id string) todo.Todo
	Check(todo todo.Todo)
	Create(todo todo.Todo)
	DelChecked()
}

func NewRepo() Repo {
	return &repo{}
}

func (*repo) All() []todo.Todo {
	var todos []todo.Todo
	firebase.All(col, &todos)
	return todos
}

func (*repo) ByID(id string) todo.Todo {
	var t todo.Todo
	firebase.ById(col, id, &t)
	return t
}

func (*repo) Check(todo todo.Todo) {
	err := firebase.Update(col, todo.ID, map[string]interface{}{
		"is_checked": todo.IsChecked,
	})
	if err != nil {
		log.Println(err)
	}
}

func (*repo) Create(todo todo.Todo) {
	id := datetime.Now(datetime.DEFAULT)
	err := firebase.Create(col, map[string]interface{}{
		"id":         id,
		"content":    todo.Content,
		"is_checked": false,
	})
	if err != nil {
		log.Println("failed to add record to firestore:", err)
	}
}

func (*repo) DelChecked() {
	err := firebase.Delete(col, "is_checked", "==", true)
	if err != nil {
		log.Println(err)
	}
}
