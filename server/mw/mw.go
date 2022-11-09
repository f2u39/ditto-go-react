// Package middleware
package mw

import (
	"crypto/rand"
	db_redis "ditto/db/redis"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var expiration = time.Duration(10 * time.Hour) // 10 hours

func SetAuth(c *gin.Context, userID string) string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("Failed to generate a random value...")
	}
	token := base64.URLEncoding.EncodeToString(b)

	log.Println("setAuth:", token, "→", userID)
	if err := db_redis.Set(token, userID, expiration); err != nil {
		panic("Failed to set session key to Redis..." + err.Error())
	}
	c.SetCookie("auth_token", token, 0, "/", "", false, false)
	return token
}

func getAuth(c *gin.Context) interface{} {
	authToken, _ := c.Cookie("auth_token")
	userID, err := db_redis.Get(authToken)

	switch {
	case err == redis.Nil:
		fmt.Println("Cannot not found user from this token...")
		return nil
	case err != nil:
		fmt.Println("Error occurred", err.Error())
		return nil
	}
	return userID
}

func ClearAuth(c *gin.Context) {
	authToken, _ := c.Cookie("auth_token")
	db_redis.Del(authToken)
	c.SetCookie("auth_token", "", -1, "/", "", false, false)
}

// Check user authority
func Auth(c *gin.Context) {
	uid := getAuth(c)
	if uid == nil {
		c.Abort()
		return
	}

	c.Next()
}
