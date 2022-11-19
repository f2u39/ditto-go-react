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
	incGroup := e.Group("/api/inc").Use(mw.Auth)
	{
		incGroup.GET("/", index)
		incGroup.POST("/create", create)
		incGroup.Any("/update", update)
	}
}

func index(c *gin.Context) {
	incs := h.IncService.All()

	data := gin.H{
		"incs": incs,
	}
	c.JSON(http.StatusOK, data)
}

func create(c *gin.Context) {
	name := c.PostForm("name")
	isDev, _ := strconv.Atoi(c.PostForm("is_developer"))
	isPub, _ := strconv.Atoi(c.PostForm("is_publisher"))

	if h.IncService.IsExists(name) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": name + " is exists!"})
		return
	} else {
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
