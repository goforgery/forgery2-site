# Guide

* __[Getting Started](#Getting_Started)__
* __[Google App Engine](#Google_App_Engine)__

## <a class="jump" name="Getting_Started"></a>Getting Started

With golang installed ([download](http://golang.org/doc/install)), get your first application started by installing __forgery__.

    go get github.com/ricallinson/forgery

Now create the file `init.go` with the following content.

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

[Source code for this example](https://github.com/ricallinson/forgery-site/tree/master/examples/local)

## <a class="jump" name="Google_App_Engine"></a>Google App Engine

You must have __forgery__ and the [Google App Engine Go Runtime](https://developers.google.com/appengine/docs/go/) installed.

In a new directory create the file `init.go` with the following content.

    package main

    import(
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

Now create the Google App Engine configuration file `app.yaml` with the following content.

    application: helloworld
    version: 1
    runtime: go
    api_version: go1

    handlers:
    - url: /.*
      script: _go_app 

With these 2 files complete start the local Google App Engine.

    /path/to/go_appengine/dev_appserver.py .

Now you can view the page in a browser [http://localhost:8080/](http://localhost:8080/)

[Source code for this example](https://github.com/ricallinson/forgery-site/tree/master/examples/googleappengine)
