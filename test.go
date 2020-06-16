package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

var db *sql.DB

func main() {

	var config struct {
		PostgresDB struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Dbname   string `yaml:"dbname"`
		} `yaml:"postgresdb"`
	}

	secret, err := ioutil.ReadFile("secret.yaml")
	if err != nil {
		log.Fatalln("Failed to load secret.yaml")
	}

	err = yaml.Unmarshal([]byte(secret), &config)
	if err != nil {
		log.Fatalf("cannot unmarshal secret.yaml: %v", err)
	}

	dbconf := config.PostgresDB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbconf.Host, dbconf.Port, dbconf.User, dbconf.Password, dbconf.Dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	defer db.Close()

	rows, err := db.Query("SELECT email, password FROM users")

	for rows.Next() {
		var email string
		var password string
		err = rows.Scan(&email, &password)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v %v\n", email, password)
	}
}
