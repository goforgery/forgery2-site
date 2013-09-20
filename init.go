package main

import(
    "net/http"
    "github.com/ricallinson/forgery"
)

func init() {
    app := forgery.CreateServer()
    app.Get("/", func(req *forgery.Request, res *forgery.Response, next func()) {
        res.End("Forgery")
    })
    http.Handle("/", app)
}