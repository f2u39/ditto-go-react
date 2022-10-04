package inc

import (
	h "ditto/ctrl"
	"ditto/model/inc"
	"ditto/mw"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	incGroup := e.Group("/inc").Use(mw.Auth)
	{
		incGroup.GET("/", index)
		incGroup.Any("/create", create)
		incGroup.Any("/update", update)
	}
}

func index(c *gin.Context) {
	incs := h.IncService.All()

	data := gin.H{
		"incs": incs,
	}
	h.RESP(c, http.StatusOK, "inc/index", data)
}

func create(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "inc/create", gin.H{})

	case "POST":
		name := c.PostForm("name")
		isDev, _ := strconv.Atoi(c.PostForm("is_developer"))
		isPub, _ := strconv.Atoi(c.PostForm("is_publisher"))

		h.IncService.Create(inc.Inc{
			Name:        name,
			IsDeveloper: isDev,
			IsPublisher: isPub,
		})
		c.Redirect(http.StatusSeeOther, "/game")
	}
}

func update(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id := c.Query("id")
		inc := h.IncService.ByID(id)

		c.HTML(http.StatusOK, "inc/update", gin.H{
			"inc": inc,
		})

	case "POST":
		id := c.PostForm("id")
		name := c.PostForm("name")
		isDev, _ := strconv.Atoi(c.PostForm("is_developer"))
		isPub, _ := strconv.Atoi(c.PostForm("is_publisher"))

		inc := h.IncService.ByID(id)
		inc.Name = name
		inc.IsDeveloper = isDev
		inc.IsPublisher = isPub
		h.IncService.Update(inc)
		c.Redirect(http.StatusSeeOther, "/")
	}
}
