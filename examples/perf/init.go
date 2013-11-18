/*
    To run this you'll need to have golang installed.

    go get github.com/ricallinson/forgery
    go run -a init.go
    curl -i http://127.0.0.1:3000/
*/

package main

import(
    "runtime"
    "github.com/ricallinson/forgery"
)

func main() {
    runtime.GOMAXPROCS(1)
    app := f.CreateServer()
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Set("Connection", "keep-alive")
        res.Send("Hello World")
    })
    app.Listen(3000)
}