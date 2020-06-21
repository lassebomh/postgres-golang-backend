# postgres-golang-backend
A backend template that combines PostgreSQL and Golang. Here is an example of how the route is written:
```
{
    switch dir() {
    case "GET":                                 // GET www.example.com
        send(200, "index")
    case "user":
        switch dir() {
        case "GET":
            send(200, "get user")               // GET www.example.com/user
        case "POST":
            send(200, "create new user")        // POST www.example.com/user
        }
    }
}
```
The purpose of this format is to create an easy and extensible way of making REST APIs in Go.
## How to use
1. Rename the `secret.yaml --example` file to `secret.yaml`. **This file contains confedential information such as the postgres user password** and other miscellaneous settings. **It is vital that this file is in your `.gitingore`!**
2. Locate `// The Route` inside `main.go`. Everything within the following curlies is where your write your API.
Have fun!
