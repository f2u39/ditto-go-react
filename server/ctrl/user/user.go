package user

import (
	h "ditto/ctrl"
	"ditto/model/user"
	"ditto/mw"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	e.Any("/user/login", login)
	e.Any("/api/user/login", loginApi)
	e.Use(mw.Auth).Any("/user/logout", logout)
}

// Handle user login
func login(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "user/login", gin.H{})

	case "POST":
		username := c.PostForm("username")
		password := c.PostForm("password")

		u, ok := h.UserService.Login(username, password)
		if ok {
			mw.SetAuth(c, u.ID.Hex())
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Login failed.",
			})
		}
	}
}

// Handle api user login
func loginApi(c *gin.Context) {
	var u user.User
	c.BindJSON(&u)

	u, ok := h.UserService.Login(u.Username, u.Password)
	if ok {
		token := mw.SetAuth(c, u.ID.Hex())
		c.JSON(http.StatusOK, gin.H{
			"auth_token": token,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Login failed.",
		})
	}
}

// Handle user logout
func logout(c *gin.Context) {
	mw.ClearAuth(c)
	c.Redirect(http.StatusSeeOther, "/user/login")
}
