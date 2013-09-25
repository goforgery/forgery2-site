package main

import(
    "net/http"
    "github.com/ricallinson/forgery"
    "github.com/ricallinson/fmustache"
    "github.com/ricallinson/fmarkdown"
)

func init() {

    app := f.CreateServer()

    app.Locals["title"] = "Forgery"

    app.Engine(".html", fmustache.Make())

    app.Get("/api.html", func(req *f.Request, res *f.Response, next func()) {
        res.Render("index.html", map[string]string{
            "title": "API Reference - Forgery",
            "body": fmarkdown.Render("./en/api.md"),
        })
    })

    /*
        Default Page.
    */
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Render("index.html", map[string]string{
            "body": fmarkdown.Render("./en/home.md"),
        })
    })

    http.Handle("/", app)
}