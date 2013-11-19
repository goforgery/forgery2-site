package main

import (
    "github.com/ricallinson/forgery"
    "time"
)

func main() {
    app := f.CreateServer()
    app.Use(f.ResponseTime())
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        time.Sleep(100 * time.Millisecond)
        res.Send("Response Time")
    })
    app.Listen(3000)
}
