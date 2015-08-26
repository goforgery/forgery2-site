package main

import (
	"errors"
	"github.com/ricallinson/forgery"
)

func main() {
	app := f.CreateServer()
	app.Use(f.ErrorHandler("Error Handler"))
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		res.Error = errors.New("Error!!!")
	})
	app.Get("/panic", func(req *f.Request, res *f.Response, next func()) {
		panic("panic!!!")
	})
	app.Listen(3000)
}
