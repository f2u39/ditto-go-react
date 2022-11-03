package main

import (
	"ditto/model/user"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	col = "user"
)

func main() {
	http.HandleFunc("/", RegisterGET)
	http.HandleFunc("/user/register", RegisterPOST)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// Display user register page
func RegisterGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("register.html"))
	tmpl.Execute(w, nil)
}

// Handle user register
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	_, err := Register(username, password)
	if err != nil {
		w.Write([]byte("Failed!"))
	} else {
		w.Write([]byte("Successed!"))
	}
}

// Handle user register form submission
func Register(username string, password string) (user.User, error) {
	_, err := ByUsername(username)

	if err == mongo.ErrNotFound {
		u := user.User{}
		u.ID = bson.NewObjectId()
		u.Username = username

		// u.Username = strings.Split(email, "@")[0]
		u.Password = HashPass(password)
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()

		dialInfo := &mgo.DialInfo{
			Addrs:    []string{"127.0.0.1"},
			Database: "ditto",
		}
		mgo, err := mgo.DialWithInfo(dialInfo)
		if err != nil {
			panic(err)
		}
		defer mgo.Close()

		db := mgo.DB("")
		c := db.C(col)

		if c.Insert(u) != nil {
			log.Println(err)
			return user.User{}, err
		}

		fmt.Println("Done.", u)
		return u, err
	}
	return user.User{}, err
}

// Return a hashed password
func HashPass(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}

// Get user by username
func ByUsername(username string) (user.User, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1"},
		Database: "ditto",
	}
	mgo, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer mgo.Close()
	filter := bson.M{"username": username}

	db := mgo.DB("")
	c := db.C(col)

	var findOne user.User
	err2 := c.Find(filter).One(&findOne)
	if err2 != nil {
		log.Println(err2)
	}
	return findOne, err2
}
