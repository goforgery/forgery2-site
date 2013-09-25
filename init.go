package main

import(
    "net/http"
    "github.com/ricallinson/forgery"
    "github.com/ricallinson/fmustache"
    "github.com/ricallinson/fmarkdown"
)

func init() {

    app := f.CreateServer()

    app.Engine(".html", fmustache.Make())

    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Render("index.html", map[string]string{
            "title": "Forgery",
            "body": fmarkdown.Render("./en/home.md"),
        })
    })
    http.Handle("/", app)
}