package word

import (
	h "ditto/ctrl"
	"ditto/model/word"
	"ditto/mw"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	e.Use(static.Serve("/word", static.LocalFile("./web", true)))
	word := e.Group("/word").Use(mw.Auth)
	{
		word.GET("/", index)
		word.Any("/create", create)
		word.POST("/check", check)
		word.Any("/update", update)
		word.GET("/delete", delete)
	}

	api := e.Group("/api/word")
	{
		api.GET("/", index)
	}
}

func check(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		id := c.PostForm("id")
		h.WordService.Check(1, id)
	}
}

func create(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "word/create", gin.H{})
	case "POST":
		w := word.Word{
			Word:    strings.TrimSpace(c.PostForm("word")),
			Example: c.PostForm("example"),
			Meaning: c.PostForm("meaning"),
		}
		h.WordService.Create(w)
		c.Redirect(http.StatusSeeOther, "/word")
	}
}

func index(c *gin.Context) {
	isChecked, _ := strconv.Atoi(c.Query("is_checked"))
	data := gin.H{
		"words": h.WordService.ByIsChecked(isChecked),
	}
	c.JSON(http.StatusOK, data)
}

func update(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id := c.Query("id")
		w := h.WordService.ByID(id)
		c.HTML(http.StatusOK, "word/update", gin.H{
			"word": w,
		})

	case "POST":
		id := c.PostForm("id")
		word := strings.TrimSpace(c.PostForm("word"))
		example := c.PostForm("example")
		meaning := c.PostForm("meaning")

		w := h.WordService.ByID(id)
		w.Word = word
		w.Example = example
		w.Meaning = meaning
		h.WordService.Update(w)
		c.Redirect(http.StatusSeeOther, "/word")
	}
}

func delete(c *gin.Context) {
	id := c.Query("id")
	h.WordService.Delete(id)
	c.Redirect(http.StatusSeeOther, "/word")
}
