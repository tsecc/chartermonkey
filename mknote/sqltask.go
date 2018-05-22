package mknote

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

//InitDB initialize a DB object
func InitDB() {
	host := os.Getenv("DB_URL")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected")
	}
	//defer db.Close()
}

//Add the profile on the list
func Add(profileName string) int64 {
	//STEP 1: check for duplication
	//STEP 2: update reservation set data = jsonb_set(data, '{name_list, 999999}', '"JH"', TRUE) where data->>'date'='2018-05-31';
	//	addQuery := "update reservation set data = jsonb_set(data, '{name_list, 999999}', '\"" + name.Displayname + "\"', TRUE) where data->>'date'='2018-05-31'"
	addQuery := "update reservation set data = jsonb_set(data, '{name_list, 999999}', '\"" + profileName + "\"', TRUE) where data->>'date'='2018-05-17'"

	result, err := db.Exec(addQuery)
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rowCount)

	return rowCount
}

func del(profile string) {

}

//Query queries specific week for the attendees
func Query() string {
	query := `SELECT data FROM reservation WHERE data @> '{"date": "2018-05-17"}'`
	var result string

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			panic(err)
		}
	}

	return result
}
