package core

import (
	handler "ditto/ctrl"
	router "ditto/ctrl/router"
	"ditto/db"
	db_redis "ditto/db/redis"
	"ditto/model/config"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Ditto ← Core
var Ditto Core

type Core struct {
	Engine *gin.Engine
	// Redis   *redis.Client
}

// Init initialize components
func Init() {
	// #1 Read app.json
	config.NewConfig()

	// #2 Initialize core
	NewCore()

	// Connect to MongoDB
	db.Init()
}

func NewCore() {
	Ditto = Core{}

	// #1 New Gin engine
	Ditto.Engine = NewEngine()

	// #2 Connect to Redis
	// Ditto.Redis = db_redis.NewRedisClient()
	db_redis.Cli = db_redis.NewRedisClient()

	// #3 Set routers
	router.Route(Ditto.Engine)

	// #4 Initialize handler
	handler.Init()

	// #5 Set logger
	Ditto.Engine.Use(gin.Logger())
}

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.Static("/asset", "./asset")
	r.Static("/assets", "../assets")
	r.LoadHTMLGlob("views/**/*")

	// CORS
	r.Use(cors.Default())
	return r
}

// Write log to a file
func SetLog() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}
