package mknote

import (
	"database/sql"
)

var db *sql.DB

func add(profile string) {

}

func del(profile string) {

}

//Query queries specific week for the attendees
func Query() string {
	query := `SELECT data FROM reservation WHERE data @> '{"date": "2018-05-24"}'`
	var result string
	err := db.QueryRow(query).Scan(&result)
	if err != nil {
		panic(err)
	}
	return result
}
