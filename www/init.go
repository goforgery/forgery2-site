package main

import (
	"github.com/goforgery/favicon"
	"github.com/goforgery/forgery2"
	"github.com/goforgery/responsetime"
	"github.com/goforgery/static"
	"github.com/goforgery/markdown"
	"github.com/ricallinson/fmustache"
	"net/http"
)

func init() {

	app := f.CreateApp()

	app.Use(responsetime.Create())
	app.Use(favicon.Create())
	app.Use(static.Create())

	app.Engine(".html", fmustache.Make())

	app.Locals["title"] = "forgery - web application framework for golang"

	// API Reference Page.
	app.Get("/api.html", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "API Reference - Forgery"
		res.Render("index.html", map[string]string{
			"body":  markdown.Render("./en/api.md"),
			"class": "index",
		})
	})

	// Guide Page.
	app.Get("/guide.html", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "API Reference - Forgery"
		res.Render("index.html", map[string]string{
			"body":  markdown.Render("./en/guide.md"),
			"class": "index",
		})
	})

	// Default Page.
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Render("index.html", map[string]string{
			"body": markdown.Render("./en/home.md"),
		})
	})

	http.Handle("/", app)
}
