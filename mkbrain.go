package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
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
					if message.Text == "恰特猴" {
						message.Text = "幹嘛~?"
					} else if message.Text == "本週+1" {
						profile, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
						if err != nil {
							log.Print(err)
						}
						message.Text = "好喔, 本週 " + profile.DisplayName + " +1, 吱吱"
					}
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
					if err != nil {
						log.Print(err)
					}
				}

			}
		}
	})
}
