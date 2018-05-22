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
	if message == "恰特猴" {
		reply = "幹嘛~?"
	} else if message == "list" {
		reply = mknote.Query()
	} else if message == "+1" && event.Source.GroupID != "" {
		profile, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
		if err != nil {
			log.Print(err)
		}
		resultBool := mknote.Add(profile.DisplayName) //returned a in64 1=added, 0=failed
		tmpl, err := template.ParseFiles("message.tpl")
		if err != nil {
			panic(err)
		}

		tmplname := "failed" //default to failed if SQL update has failed
		if resultBool == 1 {
			tmplname = "plusone"
		}

		var bytedata bytes.Buffer
		tmpl.ExecuteTemplate(&bytedata, tmplname, profile.DisplayName)
		if err != nil {
			panic(err)
		}
		reply = bytedata.String()
		//reply = "好喔, " + profile.DisplayName + " +1, 吱吱"
	} else if message == "+1" && event.Source.GroupID == "" {
		profile, err := bot.GetProfile(event.Source.UserID).Do()
		if err != nil {
			log.Print(err)
		}
		result := mknote.Add(profile.DisplayName)
		log.Printf("%d", result)
		//date := time.Now().Local().Format("2014-07-07")
		reply = "好喔, " + profile.DisplayName + " +1, 吱吱"
	} else {
		reply = "吱吱, 我聽不懂哦"
	}

	return reply
}
