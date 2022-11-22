package router

import (
	"ditto/ctrl/act"
	"ditto/ctrl/game"
	"ditto/ctrl/inc"
	"ditto/ctrl/user"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	act.Route(e)
	game.Route(e)
	inc.Route(e)
	user.Route(e)
}
