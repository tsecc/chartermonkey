package main

import (
	"bytes"
	"chartermonkey/mknote"
	"html/template"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

//Ideally, this function should query a DB, get a good answer as the reply
func reply(message string, event *linebot.Event, bot *linebot.Client) (reply string) {
	switch message {
	case "恰特猴":
		reply = "幹嘛~?" //need to use tmpl
	case "list":
		reply = mknote.Query()
	case "+1":
		if event.Source.GroupID != "" {
			//group add
			profile, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
			if err != nil {
				log.Print(err)
			}

			resultBool := mknote.Add(profile.DisplayName) //returned a in64 1=added, 0=failed
			tplID := "failed"                             //default to failed if SQL update has failed
			if resultBool == 1 {
				tplID = "plusone"
			}
			log.Print(tplID)

			name := Profile{profile.DisplayName}
			reply = assembleReply(tplID, name)
			log.Print(reply)
		} else {
			//personal add, shouldn't happen...only for admin.
			reply = "建議不要私底下揪團~吱吱" //need to use tmpl
		}
	}

	return reply
}

func assembleReply(tplID string, myProfile Profile) string {
	var bytedata bytes.Buffer
	tmpl, err := template.ParseFiles("message.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(&bytedata, tplID, myProfile)
	if err != nil {
		panic(err)
	}
	reply := bytedata.String()
	return reply
}
