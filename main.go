package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func popAndCheck(xs *[]string, path string) bool {

	if (*xs)[0] == path {
		*xs = append((*xs)[1:])
		return true
	} else {
		return false
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
			return popAndCheck(&r, s)
		}

		send := func(statusCode int, body string) {
			res.WriteHeader(statusCode)
			io.WriteString(res, body)
		}

		{
			res.WriteHeader(404)

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
