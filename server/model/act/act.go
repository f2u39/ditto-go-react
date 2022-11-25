package act

import (
	"ditto/model/game"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	GAMING      = Type("Gaming")
	PROGRAMMING = Type("Programming")
)

type Act struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	StartTime time.Time          `json:"start" bson:"start"`
	EndTime   time.Time          `json:"end" bson:"end"`
	Date      string             `json:"date" bson:"date"`
	Duration  int                `json:"duration" bson:"duration"`
	Type      Type               `json:"type" bson:"type"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`

	GameID primitive.ObjectID `json:"game_id" bson:"game_id,omitempty"`
}

type Detail struct {
	Act  Act         `json:"act" bson:",inline"`
	Game []game.Game `json:"game" bson:"game"`

	Hour int `json:"hour"`
	Min  int `json:"min"`
}

type Summary struct {
	GameDur  int `json:"game_dur"`
	GameHour int `json:"game_hour"`
	GameMin  int `json:"game_min"`

	PgmDur  int `json:"pgm_dur"`
	PgmHour int `json:"pgm_hour"`
	PgmMin  int `json:"pgm_min"`
}

type Type string

type StopWatch struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  int       `json:"duration"`
	Type      Type      `json:"type"`
	GameID    string    `json:"game_id"`
	GameTitle string    `json:"game_title"`
}

func NewStopWatch() *StopWatch {
	return &StopWatch{time.Time{}, time.Time{}, 0, "", "", ""}
}

func (sw *StopWatch) Start(typ, gid, gtl string) {
	sw.Type = Type(typ)
	sw.StartTime = time.Now()
	sw.GameID = gid
	sw.GameTitle = gtl
}

// / Set *stopwatch to nil and return end time
func (sw *StopWatch) Stop() time.Time {
	sw.EndTime = time.Now()
	sw.Duration = int(sw.EndTime.Sub(sw.StartTime).Minutes())
	endTime := sw.EndTime
	return endTime
}

func (sw *StopWatch) Reset() {
	sw = nil
}
