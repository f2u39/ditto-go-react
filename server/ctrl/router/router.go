package router

import (
	"ditto/ctrl/act"
	"ditto/ctrl/game"
	"ditto/ctrl/home"
	"ditto/ctrl/inc"
	"ditto/ctrl/line"

	// "ditto/ctrl/todo"
	"ditto/ctrl/user"
	"ditto/ctrl/word"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	act.Route(e)
	game.Route(e)
	home.Route(e)
	inc.Route(e)
	line.Route(e)
	// todo.Route(e)
	user.Route(e)
	word.Route(e)
}
