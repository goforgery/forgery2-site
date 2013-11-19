package main

import (
    "github.com/ricallinson/forgery"
)

func main() {
    app := f.CreateServer()
    app.Use(f.ResponseTime())
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Send("Response Time")
    })
    app.Listen(3000)
}
