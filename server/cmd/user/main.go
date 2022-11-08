package main

import (
	"context"
	"ditto/model/user"
	"log"
	"net/http"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	MGO_DB  = "ditto"
	MGO_COL = "user"

	cli *mongo.Client
	col *mongo.Collection
)

func main() {
	http.HandleFunc("/", GET)
	http.HandleFunc("/user/register", POST)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Display user register page
func GET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("register.html"))
	tmpl.Execute(w, nil)
}

// Handle user register
func POST(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := register(username, password)
	if err != nil {
		w.Write([]byte("failed!"))
	} else {
		w.Write([]byte("successed!"))
	}
}

// Handle user register form submission
func register(username string, password string) error {
	u, err := byUsername(username)

	if err != nil && len(u.Username) == 0 {
		new := user.User{}
		new.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		new.Username = username
		new.Password = HashPass(password)
		new.CreatedAt = time.Now()
		new.UpdatedAt = time.Now()

		connect()
		defer disconnect()

		_, err := col.InsertOne(context.TODO(), new)
		if err != nil {
			return err
		}
	}

	return nil
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

// Connect to MongoDB
func connect() error {
	var err error
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	cli, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = cli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	col = cli.Database(MGO_DB).Collection(MGO_COL)

	return nil
}

func disconnect() {
	cli.Disconnect(context.TODO())
}

// Get user by username
func byUsername(username string) (user.User, error) {
	connect()
	defer disconnect()

	var u user.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	err := col.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, err
}
