package game

import (
	"ditto/model/inc"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string
type Genre string
type Platform string

// Game model
type Game struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Genre       string             `json:"genre" bson:"genre"`
	Platform    Platform           `json:"platform" bson:"platform"`
	DeveloperID primitive.ObjectID `json:"developer_id" bson:"developer_id"`
	PublisherID primitive.ObjectID `json:"publisher_id" bson:"publisher_id"`
	Status      Status             `json:"status" bson:"status"`
	PlayTime    int                `json:"playtime" bson:"play_time"`
	Ranking     int                `json:"ranking" bson:"ranking"`
	Rating      string             `json:"rating" bson:"rating"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type Detail struct {
	Game      Game    `json:"game" bson:",inline"`
	Developer inc.Inc `json:"developer" bson:"developer"`
	Publisher inc.Inc `json:"publisher" bson:"publisher"`
	PlayHour  int     `json:"play_hour"`
	PlayMin   int     `json:"play_min"`
}

// Return genres
func Genres() []Genre {
	return []Genre{
		ACT, ARPG, CARD, FPS, MOBA, RHYTHM, RPG, RTS,
		SB, SIMULATION, SURVIVAL, TPS, VN,
	}
}

// Return platforms
func Platforms() []Platform {
	return []Platform{
		PC, PLAYSTATION, SWITCH, MOBILE,
	}
}

// Return status
func Statuses() []Status {
	return []Status{
		PLAYING, PLAYED, BLOCKING,
	}
}

var (
	// Status
	PLAYING  = Status("Playing")
	PLAYED   = Status("Played")
	BLOCKING = Status("Blocking")

	// Platform
	PC          = Platform("PC")
	SWITCH      = Platform("NintendoSwitch")
	PLAYSTATION = Platform("PlayStation")
	XBOX        = Platform("Xbox")
	MOBILE      = Platform("Mobile")

	// Genres
	ACT        = Genre("ACT")
	ARPG       = Genre("ARPG")
	CARD       = Genre("Card")
	FPS        = Genre("FPS")
	MOBA       = Genre("MOBA")
	RHYTHM     = Genre("Rhythm")
	RPG        = Genre("RPG")
	RTS        = Genre("RTS")
	SB         = Genre("Sandbox")
	SIMULATION = Genre("Simulation")
	SURVIVAL   = Genre("Survival")
	TPS        = Genre("TPS")
	VN         = Genre("Visual Novel")
)
