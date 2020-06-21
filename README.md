# postgres-golang-backend
A backend template that combines PostgreSQL and Golang.
The route has the following structure:
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
1. Rename the `secret.yaml --example` to `secret.yaml`. **This contains confedential information such as the postgres user password** and other miscellaneous settings. **It is vital that this file is in your `.gitingore`** for obvious reasons.
2. Locate *the route* inside `main.go`. This is where the you write the API.
