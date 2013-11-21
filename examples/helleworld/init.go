package main

import (
	"github.com/ricallinson/forgery"
)

func main() {
	app := f.CreateServer()
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Send("Hello world.")
	})
	app.Listen(3000)
}
