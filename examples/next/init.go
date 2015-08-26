package main

import (
	"github.com/ricallinson/forgery"
	"github.com/ricallinson/stackr"
)

func main() {
	app := f.CreateServer()
	app.Use(func(req *stackr.Request, res *stackr.Response, next func()) {
		res.Write("Front\n")
		next()
		res.Write("Back\n")
	})
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Write("Middle\n")
	})
	app.Listen(3000)
}
