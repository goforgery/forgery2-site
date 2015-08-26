package main

import (
	"github.com/goforgery/favicon"
	"github.com/goforgery/forgery2"
	"github.com/goforgery/static"
	"github.com/ricallinson/fmarkdown"
	"github.com/ricallinson/fmustache"
	"net/http"
)

func init() {

	app := f.CreateApp()

	// app.Use(f.ResponseTime())
	app.Use(favicon.Create())
	app.Use(static.Create())

	app.Engine(".html", fmustache.Make())

	app.Locals["title"] = "forgery - web application framework for golang"

	// API Reference Page.
	app.Get("/api.html", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "API Reference - Forgery"
		res.Render("index.html", map[string]string{
			"body":  fmarkdown.Render("./en/api.md"),
			"class": "index",
		})
	})

	// Guide Page.
	app.Get("/guide.html", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "API Reference - Forgery"
		res.Render("index.html", map[string]string{
			"body":  fmarkdown.Render("./en/guide.md"),
			"class": "index",
		})
	})

	// Default Page.
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Render("index.html", map[string]string{
			"body": fmarkdown.Render("./en/home.md"),
		})
	})

	http.Handle("/", app)
}
