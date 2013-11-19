package main

import (
	"github.com/ricallinson/forgery"
	"net/http"
)

func init() {
	app := f.CreateServer()
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Send("Hello world.")
	})
	http.Handle("/", app)
}
