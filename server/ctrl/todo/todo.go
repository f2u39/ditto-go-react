package todo

import (
	h "ditto/ctrl"
	"ditto/model/todo"
	"ditto/mw"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	todo := e.Group("/todo").Use(mw.Auth)
	{
		todo.GET("/", index)
		todo.POST("/create", create)
		todo.POST("/check", check)
	}
}

func index(c *gin.Context) {
	h.TodoService.DelChecked()
	todos := h.TodoService.All()
	c.HTML(http.StatusOK, "todo/index", gin.H{
		"todos": todos,
	})
}

func create(c *gin.Context) {
	h.TodoService.Create(todo.Todo{
		Content: c.PostForm("content"),
	})
	c.Redirect(http.StatusSeeOther, "/todo")
}

func check(c *gin.Context) {
	id := c.PostForm("id")
	isChecked, _ := strconv.Atoi(c.PostForm("is_checked"))
	h.TodoService.Check(id, isChecked)
}
