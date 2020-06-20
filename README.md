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
