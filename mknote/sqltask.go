package mknote

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

//ResultSet stores result
type ResultSet struct {
	Namelist string
}

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
func Add(profileName string, date string) int64 {
	//STEP 1: check for duplication
	addQuery := `UPDATE reservation SET data = jsonb_set(data, '{name_list, 999999}', '"` + profileName + `"', TRUE) WHERE data->>'date'='` + date + `';`

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
	//don't use UPDATE reservation SET data=data - '{name_list, "FreddyChen"}' WHERE data @> '{"date": "2018-06-07"}';
	//start from 0, UPDATE reservation SET data=data #- '{name_list, 1}' WHERE data @> '{"date": "2018-06-07"}';
}

//Query queries specific week for the attendees
func Query(date string) []ResultSet {
	//date = "2018-06-07"
	query := `SELECT jsonb_array_elements_text(data->'name_list') as name FROM reservation WHERE data @> '{"date": "` + date + `"}';`
	var result string
	resultSet := []ResultSet{}

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
		tmp := ResultSet{result}
		resultSet = append(resultSet, tmp)
	}
	return resultSet
}
