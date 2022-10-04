package handler

import (
	act_srv "ditto/service/act"
	game_srv "ditto/service/game"
	inc_srv "ditto/service/inc"
	todo_srv "ditto/service/todo"
	user_srv "ditto/service/user"
	word_srv "ditto/service/word"

	"github.com/gin-gonic/gin"
)

var (
	ActService  act_srv.Service
	GameService game_srv.GameService
	IncService  inc_srv.IncService
	TodoService todo_srv.Service
	UserService user_srv.UserService
	WordService word_srv.WordService
)

func Init() {
	ActService = act_srv.NewService()
	GameService = game_srv.NewService()
	IncService = inc_srv.NewIncService()
	TodoService = todo_srv.NewService()
	UserService = user_srv.NewUserService()
	WordService = word_srv.NewWordService()
}

func RESP(c *gin.Context, code int, name string, obj interface{}) {
	path := c.Request.URL.Path

	if len(path) > 4 && path[0:4] == "/api" {
		c.JSON(code, obj)
		return
	}

	c.HTML(code, name, obj)
}
