package main

import (
    "github.com/ricallinson/forgery"
)

func main() {
    app := f.CreateServer()
    app.Use(f.Logger())
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Send("Logging")
    })
    app.Listen(3000)
}
