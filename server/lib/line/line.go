package line

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	Bot *linebot.Client
)

func NewLineBot() {
	var err error
	Bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
