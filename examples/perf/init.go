package main

import(
    "runtime"
    "github.com/ricallinson/forgery"
)

func main() {
    runtime.GOMAXPROCS(1)
    app := f.CreateServer()
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Set("Connection", "keep-alive")
        res.Send("Hello World")
    })
    app.Listen(3000)
}