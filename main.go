package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

var db *sql.DB

var config struct {
	PostgresDB struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Dbname   string `yaml:"DBname"`
	} `yaml:"PostgresDB"`
	HTTP struct {
		Port int `yaml:"Port"`
	} `yaml:"HTTP"`
}

func main() {

	secretFile, err := ioutil.ReadFile("secret.yaml")
	if err != nil {
		log.Fatalln("Failed to load secret.yaml")
	}

	err = yaml.Unmarshal([]byte(secretFile), &config)
	if err != nil {
		log.Fatalf("cannot unmarshal secret.yaml: %v", err)
	}

	dbconf := config.PostgresDB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconf.Host, dbconf.Port, dbconf.User, dbconf.Password, dbconf.Dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	defer db.Close()

	// rows, err := db.Query("SELECT email, password FROM users")
	// for rows.Next() {
	// 	var email string
	// 	var password string
	// 	err = rows.Scan(&email, &password)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%v %v\n", email, password)
	// }

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				statusCode := err.(int)
				fmt.Printf("%T %v", err, err)
				res.WriteHeader(statusCode)
			}
		}()

		r := strings.SplitN(strings.Trim(req.URL.Path, "/")+"/"+req.Method, "/", -1)

		if r[0] == "" {
			r = append(r[:0], r[1:]...)
		}

		dir := func() string {
			return func(xs *[]string) string {
				next := (*xs)[0]
				*xs = append((*xs)[1:])
				return next
			}(&r)
		}

		send := func(statusCode int, body string) {
			res.WriteHeader(statusCode)
			io.WriteString(res, body)
		}

		// The route
		{
			switch dir() {
			case "GET":
				send(200, "index")
			case "user":
				switch dir() {
				case "POST":
					fmt.Println()
				}
			}
		}
	})

	http.ListenAndServe(":"+strconv.Itoa(config.HTTP.Port), nil)
}
