package word

import (
	h "ditto/ctrl"
	"ditto/model/word"
	"ditto/mw"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	e.Use(static.Serve("/word", static.LocalFile("./web", true)))
	word := e.Group("/api/word").Use(mw.Auth)
	{
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
		h.WordService.Check(id, 1)
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
	// TODO Need to fix it
	// isChecked, _ := strconv.Atoi(c.Query("is_checked"))
	data := gin.H{
		// "words": h.WordService.PageIsChecked(isChecked),
	}
	c.JSON(http.StatusOK, data)
}

func update(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id := c.Query("id")
		w := h.WordService.ByID(id)
		c.JSON(http.StatusOK, gin.H{
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
