package main

import (
	"bytes"
	"chartermonkey/mknote"
	"html/template"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	dateFormat = "2006-01-02"
	sessionDay = 4 //usually play on Thursday
)

func reply(message string, event *linebot.Event, bot *linebot.Client) (reply string) {
	date := getNextSession()
	replyInfo := ReplyInfo{"", ""}
	resultset := []mknote.ResultSet{}

	switch message {
	case "恰特猴":
		replyInfo.TplID = "wazup"
		reply = assembleReply(replyInfo)
	case "list":
		resultset = mknote.Query(date)
		replyInfo.TplID = "list"
		reply = assembleList(replyInfo, resultset)
	case "+1":
		if event.Source.GroupID != "" { //+1 in group
			profile, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
			if err != nil {
				log.Print(err)
			}

			//replyInfo.Name = profile.DisplayName
			ex := checkExists(profile.DisplayName, date) //check if the name exists in list
			if ex != false {
				//run DB update and return in64, 1=added, 0=failed
				resultBool := mknote.Add(profile.DisplayName, date)
				if resultBool == 1 {
					replyInfo.TplID = "plusone"
				} else {
					replyInfo.TplID = "failed"
				}
			} else {
				replyInfo.TplID = "duplicate"
			}
		} else {
			//personal +1, shouldn't happen.
			replyInfo.TplID = "reject"
		}
		reply = assembleReply(replyInfo)
	default:
		reply = ""
	}

	return reply
}

//this func compare current day, get the next target date
func getNextSession() (targetDate string) {
	today := time.Now().AddDate(0, 0, 0)
	intOfToday := int(today.Weekday())
	// fmt.Println("Today is ", today.Format("2006-01-02"))
	// fmt.Println("Day code is", intOfToday)
	var addTo int
	if intOfToday > sessionDay {
		addTo = 4 - intOfToday + 7
	} else {
		addTo = 4 - intOfToday
	}
	targetDate = today.AddDate(0, 0, addTo).Format("2006-01-02")
	//fmt.Println("addTo is ", addTo, ",Next Session is on", targetDate)
	return targetDate
}

func checkExists(name string, date string) bool {
	//need to query for existance check
	resultset := mknote.Query(date)
	for _, a := range resultset {
		if a.Namelist == name {
			return false
		}
	}
	return true
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
func assembleList(info ReplyInfo, resultset []mknote.ResultSet) string {
	var bytedata bytes.Buffer
	tmpl, err := template.ParseFiles("message.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(&bytedata, info.TplID, resultset)
	if err != nil {
		panic(err)
	}
	reply := bytedata.String()
	return reply
}
