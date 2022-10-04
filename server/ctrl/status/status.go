package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route adds routes to router
func Route(e *gin.Engine) {
	errorGroup := e.Group("/error")
	{
		errorGroup.GET("/404", Error404)
	}
}

// Error404 displays 404 page not found
func Error404(c *gin.Context) {

}

// ShowError displays error message in error page
func ErrorMsg(c *gin.Context, msg string) {
	c.HTML(http.StatusOK, "error/error", gin.H{
		"msg": msg,
	})
}
