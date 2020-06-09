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

		r := strings.SplitN(strings.Trim(req.URL.Path, "/"), "/", -1)

		r = append(r, req.Method)

		dir := func(s string) bool { return popAndCheck(&r, s) }

		{
			res.Header().Add("Content-Type", "text/html")

			if dir("1") {
				if dir("1") {
					if dir("GET") {
						io.WriteString(res, "1 1 get")

					} else if dir("HEAD") {
						io.WriteString(res, "1 1 head")
					}
				} else if dir("2") {
					if dir("GET") {
						io.WriteString(res, "1 2 get")
					}
				} else if dir("GET") {
					io.WriteString(res, "1 get")
				}

			} else if dir("2") {
				if dir("GET") {
					io.WriteString(res, "2 get")
				}
			}
		}

	})

	http.ListenAndServe(":5000", nil)
}
