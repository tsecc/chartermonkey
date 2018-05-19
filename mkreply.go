package main

import (
	"chartermonkey/mknote"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

//Ideally, this function should query a DB, get a good answer as the reply
func reply(message string, event *linebot.Event, bot *linebot.Client) (reply string) {
	if message == "恰特猴" {
		reply = "幹嘛~?"
	} else if message == "list" {
		reply = mknote.Query()
	} else if message == "+1" && event.Source.GroupID != "" {
		profile, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
		if err != nil {
			log.Print(err)
		}
		reply = "好喔, " + profile.DisplayName + " +1, 吱吱"
	} else if message == "+1" && event.Source.GroupID == "" {
		profile, err := bot.GetProfile(event.Source.UserID).Do()
		if err != nil {
			log.Print(err)
		}
		date := time.Now().Local().Format("2014-07-07")
		reply = "好喔, 今天是" + date + ", 下次 " + profile.DisplayName + " +1, 吱吱"
	} else {
		reply = "吱吱, 我聽不懂哦"
	}

	return reply
}
