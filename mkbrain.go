package main

import (
	"chartermonkey/mknote"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	mknote.InitDB()

	mknote.Query()

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req) //events is []*Event
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events { //event is *Event
			if event.Type == linebot.EventTypeMessage { //"message" should be defined in /callback
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					message.Text = reply(message.Text, event, bot)
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
