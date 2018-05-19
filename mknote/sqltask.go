package mknote

import (
	"database/sql"
	"fmt"
	"os"
)

// const (
// 	//host     = "ec2-23-23-130-158.compute-1.amazonaws.com"
// 	//port     = 5432
// 	user     = "eqhagpctrbrinp"
// 	password = "d7251e298543dcfaab18324c787e3bc5ba0b987156d19dfe7107bee0094b3ce2"
// 	dbname   = "d59v5g8s8r2n4v"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "your-password"
// 	dbname   = "mknote"
// )

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

func add(profile string) {

}

func del(profile string) {

}

//Query queries specific week for the attendees
func Query() string {
	query := `SELECT data FROM reservation WHERE data @> '{"date": "2018-05-24"}'`
	var result string

	// err := db.QueryRow(query).Scan(&result)
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
