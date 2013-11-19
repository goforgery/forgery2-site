package main

import (
	"github.com/ricallinson/forgery"
)

func main() {
	app := f.CreateServer()
	app.Use(f.Static())
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Send("<a href=\"file.txt\">file.txt</a>")
	})
	app.Listen(3000)
}
