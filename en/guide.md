# Guide

## Getting Started

With golang installed ([download](http://golang.org/doc/install)), get your first application started by installing __forgery__.

    go get github.com/ricallinson/forgery

Now create a file `init.go` with the following content.

    package main

    import("github.com/ricallinson/forgery")

    func main() {
        app := f.CreateServer()
        app.Get("/", func(req *f.Request, res *f.Response, next func()) {
            res.Send("Hello world.")
        })
        app.Listen(3000)
    }

Start your app.

    go run init.go

Now you can view the page in a browser [http://localhost:3000/](http://localhost:3000/)