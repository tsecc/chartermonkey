package main

import (
	"bytes"
	"chartermonkey/mknote"
	"html/template"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

//ReplyInfo structured a person profile and the ID for bot's reply
type ReplyInfo struct {
	TplID string
	Name  string
}

//Ideally, this function should query a DB, get a good answer as the reply
func reply(message string, event *linebot.Event, bot *linebot.Client) (reply string) {
	replyInfo := ReplyInfo{}
	switch message {
	case "恰特猴":
		replyInfo.TplID = "wazup"
	case "list":
		replyInfo.Name = mknote.Query()
		replyInfo.TplID = "list"
	case "+1":
		if event.Source.GroupID != "" {
			//group add
			profile, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
			if err != nil {
				log.Print(err)
			}

			resultBool := mknote.Add(profile.DisplayName) //run DB update and return in64, 1=added, 0=failed
			replyInfo.Name = profile.DisplayName
			if resultBool == 1 {
				replyInfo.TplID = "plusone"
			} else {
				replyInfo.TplID = "failed"
			}
			log.Print(replyInfo.TplID)
		} else {
			//personal add, shouldn't happen...or only for admin.
			replyInfo.TplID = "reject"
		}
	default:
		reply = ""
	}
	reply = assembleReply(replyInfo)
	return reply
}

func assembleReply(info ReplyInfo) string {
	var bytedata bytes.Buffer
	tmpl, err := template.ParseFiles("message.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(&bytedata, info.TplID, info)
	if err != nil {
		panic(err)
	}
	reply := bytedata.String()
	return reply
}
