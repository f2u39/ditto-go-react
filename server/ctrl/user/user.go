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
}

// Handle user login
func login(c *gin.Context) {
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
