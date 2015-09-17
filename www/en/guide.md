# Guide

* __[Getting Started](#Getting_Started)__
* __[Serving Static Files](#Serving_Static_Files)__
* __[Google App Engine](#Google_App_Engine)__

## <a class="jump" name="Getting_Started"></a>Getting Started

With golang installed ([download](http://golang.org/doc/install)), get your first application started by installing __forgery2__.

    go get github.com/goforgery/forgery2

Now create the file `init.go` with the following content.

    package main

    import("github.com/goforgery/forgery2")

    func main() {
        app := f.CreateServer()
        app.Get("/", func(req *f.Request, res *f.Response, next func()) {
            res.Send("Hello world.")
        })
        app.Listen(3000)
    }

Start the app.

    go run init.go

Now you can view the page in a browser [http://localhost:3000/](http://localhost:3000/).

* [Source code for this example](https://github.com/goforgery/forgery2-site/tree/master/examples/helloworld)

## <a class="jump" name="Serving_Static_Files"></a>Serving Static Files

To serve static files such as CSS, JS and images you can use the __f.Static__ middleware.

Create the file `init.go` with the following content.

    package main

    import (
        "github.com/goforgery/forgery2"
    )

    func main() {
        app := f.CreateServer()
        app.Use(f.Static())
        app.Get("/", func(req *f.Request, res *f.Response, next func()) {
            res.Send("<a href=\"file.txt\">file.txt</a>")
        })
        app.Listen(3000)
    }

Now make a new directory named `public` with the a file named `file.txt` in it. The content of the file can be anything you like.

Start the app.

    go run init.go

Now you can view the page in a browser [http://localhost:3000/](http://localhost:3000/). When you click on the `file.txt` link you will be shown the content of the file.

* [Source code for this example](https://github.com/goforgery/forgery2-site/tree/master/examples/static)

## <a class="jump" name="Google_App_Engine"></a>Google App Engine

You must have __forgery2__ and the [Google App Engine Go Runtime](https://developers.google.com/appengine/docs/go/) installed.

In a new directory create the file `init.go` with the following content.

    package main

    import(
        "github.com/goforgery/forgery2"
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

With these 2 files complete start the local _Google App Engine_.

    /path/to/go_appengine/goapp serve .

Now you can view the page in a browser [http://localhost:8080/](http://localhost:8080/)

* [Source code for this example](https://github.com/goforgery/forgery2-site/tree/master/examples/googleappengine)
