package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

func checkAndPop(xs *[]string, path string) bool {

	if (*xs)[0] == path {
		*xs = append((*xs)[1:])
		return true
	} else {
		return false
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qp48shrushY!"
	dbname   = "power"
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}

func main() {
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

		dir := func(s string) bool {
			return checkAndPop(&r, s)
		}

		send := func(statusCode int, body string) {
			res.WriteHeader(statusCode)
			io.WriteString(res, body)
		}

		{
			switch {
			case dir("1"):
				switch {
				case dir("2"):
					switch {
					case dir("GET"):
						send(200, "1 2 GET")
					}
				}
			case dir("GET"):
				send(200, "index")
			}
		}
	})

	http.ListenAndServe(":5000", nil)
}
