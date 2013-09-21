package main

import(
    "net/http"
    "github.com/ricallinson/forgery"
)

func init() {
    app := f.CreateServer()
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.End("Forgery")
    })
    http.Handle("/", app)
}