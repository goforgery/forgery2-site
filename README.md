## Install

Requires you to have [Go](http://golang.org/doc/install) installed.

    go get github.com/goforgery/forgery2
    go get github.com/goforgery/markdown
    go get github.com/goforgery/mustache

## Run locally

Requires you to have the [Google App Engine Go Runtime](https://developers.google.com/appengine/docs/go/) installed.

    cd forgery2-site
    ~/Go/appengine/goapp serve ./www

## Deploy

    ~/Go/appengine/goapp deploy ./www
