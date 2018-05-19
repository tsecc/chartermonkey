package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	host     = "ec2-23-23-130-158.compute-1.amazonaws.com"
	port     = 5432
	user     = "eqhagpctrbrinp"
	password = "d7251e298543dcfaab18324c787e3bc5ba0b987156d19dfe7107bee0094b3ce2"
	dbname   = "d59v5g8s8r2n4v"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("DB connected")
	}
	defer db.Close()

	// query := `SELECT data FROM reservation WHERE data @> '{"date": "2018-05-24"}'`
	// var result string
	// err = db.QueryRow(query).Scan(&result)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(result)

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
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
