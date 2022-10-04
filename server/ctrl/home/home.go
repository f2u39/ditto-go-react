package home

import (
	"ditto/mw"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Add routes to router
func Route(e *gin.Engine) {
	home := e.Group("/").Use(mw.Auth)
	{
		home.GET("/", Index)
		home.GET("/index", Index)
	}
}

// Displays index page
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{})
}
