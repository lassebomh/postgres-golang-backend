package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "insert password here"
	dbname   = "power"
)

type User struct {
	Password   string
	Email      string
	Created_On string
}

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, _ = sql.Open("postgres", psqlInfo)
	defer db.Close()

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM users")

	for rows.Next() {
		var email string
		var password string
		var created_on string
		err = rows.Scan(&email, &password, &created_on)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v %v %v\n", email, password, created_on)
	}

}
