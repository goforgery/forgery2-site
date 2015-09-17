package main

import (
	"github.com/goforgery/favicon"
	"github.com/goforgery/forgery2"
	"github.com/goforgery/markdown"
	"github.com/goforgery/mustache"
	"github.com/goforgery/responsetime"
	"github.com/goforgery/static"
	"net/http"
)

func init() {

	app := f.CreateApp()

	app.Use(responsetime.Create())
	app.Use(favicon.Create())
	app.Use(static.Create())

	app.Engine(".html", mustache.Create())

	app.Locals["title"] = "forgery2 - web application framework for golang"

	// API Reference Page.
	app.Get("/api.html", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "API Reference - forgery2"
		res.Render("index.html", map[string]string{
			"body":  markdown.Render("./en/api.md"),
			"class": "index",
		})
	})

	// Guide Page.
	app.Get("/guide.html", func(req *f.Request, res *f.Response, next func()) {
		res.Locals["title"] = "API Reference - forgery2"
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
