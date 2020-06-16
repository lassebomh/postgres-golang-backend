package main

import (
	"database/sql"
	"encoding/json"
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

// todo: exit with status code

// todo: exit with status code and message

// todo: exit with message only (status code 500)

// todo: proper send that returns

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

				statusCode, convErr := err.(int)
				if convErr {
					fmt.Printf("Exited with status code: %v", err)
					res.WriteHeader(statusCode)
				} else {
					fmt.Printf("%T %v", err, err)
					res.WriteHeader(500)
				}
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

		// The Route
		{
			switch dir() {
			case "GET":
				send(200, "index")
			case "user":
				switch dir() {
				case "POST":
					b, _ := ioutil.ReadAll(req.Body)

					var credentialsInput struct {
						Email    string `json:"email"`
						Password string `json:"password"`
					}

					err = json.Unmarshal([]byte(b), &credentialsInput)
					if err != nil {
						log.Panicf("Failed to unmarshal input: %v", err)
					}

					var countMatches int

					err := db.QueryRow(`select count(email) from users where (email=$1)`,
						credentialsInput.Email).Scan(&countMatches)
					if err != nil {
						log.Panic("Failed to check wether email is taken")
					}

					// todo: assert valid email format

					if countMatches == 0 {
						send(200, "")
					} else {
						send(403, "")
					}
				}
			}
		}
	})

	http.ListenAndServe(":"+strconv.Itoa(config.HTTP.Port), nil)
}
