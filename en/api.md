# API Reference

## f.CreateServer()

Create a forgery application.

    package main

    import("github.com/ricallinson/forgery")

    func init() {
        app := f.CreateServer()
        app.Get("/", func(req *f.Request, res *f.Response, next func()) {
            res.Send("Hello world.")
        })
        app.Listen(3000)
    }

## Application

### app.Set(name, value)

Assigns setting `name` to `value`.

    app.Set("title", "My Site")
    app.Get("title")
    // => "My Site"

### app.Get(name)

Get setting `name` value.

    app.Get("title")
    // => ""

    app.Set("title", "My Site")
    app.Get("title")
    // => "My Site"

### app.Enable(name)

Set setting `name` to `true`.

    app.Enable("trust proxy")
    app.Get("trust proxy")
    // => "true"

### app.Disable(name)

Set setting `name` to `false`.

    app.Disable("trust proxy")
    app.Get("trust proxy")
    // => "false"

### app.Enabled(name)

Check if setting `name` is enabled.

    app.Enabled("trust proxy")
    // => false

    app.Enable("trust proxy")
    app.Enabled("trust proxy")
    // => true

### app.Disabled(name)

Check if setting `name` is disabled.

    app.Disabled("trust proxy")
    // => true

    app.Enable("trust proxy")
    app.Disabled("trust proxy")
    // => false

### app.Configure([env], callback)

Conditionally invoke `callback` when `env` matches `app.get("env")`. This method is effectively an `if` statement as illustrated in the following snippets. These functions are __not__ required in order to use `app.set()` and other configuration methods.

    // all environments
    app.Configure(func() {
        app.Set("title", "My Application")
    })

    // development only
    app.Configure("development", func() {
        app.Set("db uri", "localhost/dev")
    })

    // production only
    app.Configure("production", func() {
        app.Set("db uri", "n.n.n.n/prod")
    })

effectively sugar for:

    // all environments
    app.Set("title", "My Application")

    // development only
    if "development" == app.Get("env") {
        app.Set("db uri", "localhost/dev")
    }

    // production only
    if "production" == app.Get("env") {
        app.Set("db uri", "n.n.n.n/prod")
    }

### Application Settings

The following settings are provided to alter how Forgery will behave:

* _env_ Environment mode, defaults to `os.Getenv("GO_ENV")` or "development"
* _trust proxy_ Enables reverse proxy support, disabled by default
* TBD _jsonp callback name_ Changes the default callback name of _?callback=_
* TBD _json replacer_ JSON replacer callback, null by default
* TBD _json spaces_ JSON response spaces for formatting, defaults to _2_ in development, _0_ in production
* TBD _case sensitive routing_ Enable case sensitivity, disabled by default, treating "/Foo" and "/foo" as the same
* TBD _strict routing_ Enable strict routing, by default "/foo" and "/foo/" are treated the same by the router
* TBD _view cache_ Enables view template compilation caching, enabled in production by default
* _view engine_ The default engine extension to use when omitted
* _views_ The view directory path, defaulting to "./views"

### app.Use([path], function)

Wrapper for `stackr.Use()`. See the [Stackr](http://gostackr.appspot.com/) documentation for details.

### app.Engine(ext, callback)

Register the given template engine `callback` as `ext`.

    app.Engine(".html", fmustache.Make())

### TBD app.Param([name], callback)

This feature is not supported yet.

### app.VERB(path, [callback...], callback)

The `app.VERB()` methods provide the routing functionality in Forgery, where __VERB__ is one of the HTTP verbs, such as `app.Post()`. Multiple callbacks may be given, all are treated equally, and behave just like middleware, with the one exception that these callbacks may invoke `next("route")` to bypass the remaining route callback(s). This mechanism can be used to perform pre-conditions on a route then pass control to subsequent routes when there is no reason to proceed with the route matched.

The following snippet illustrates the most simple route definition possible. Forgery translates the path strings to regular expressions, used internally to match incoming requests. Query strings are not considered when performing these matches, for example "GET /" would match the following route, as would "GET /?name=ric".

    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Send("Hello world.")
    })

__NOTE: Regular expressions and route parameters will be supported in later releases__.

### app.All(path, [callback...], callback)

This method functions just like the `app.VERB()` methods, however it matches all HTTP verbs.

This method is extremely useful for mapping "global" logic for specific path prefixes or arbitrary matches. For example if you placed the following route at the top of all other route definitions, it would require that all routes from that point on would require authentication, and automatically load a user. Keep in mind that these callbacks do not have to act as end points, `loadUser` can perform a task, then `next()` to continue matching subsequent routes.

    app.all("*", requireAuthentication, loadUser)

Or the equivalent:

    app.all("*", requireAuthentication)
    app.all("*", loadUser)

Another great example of this is white-listed "global" functionality. Here the example is much like before, however only restricting paths prefixed with "/api":

    app.all("/api/", requireAuthentication)


### app.Locals

Application local variables are provided to all templates rendered within the application. This is useful for providing app-level data.

    app.Locals["title"] = "My App";

### app.Render(view, [locals...])

Render a `view` file returning with the rendered string. This is the app-level variant of `res.Render()`, and otherwise behaves the same way.

    s := res.Render("index.html", map[string]string{
        "body": "Document body",
    })

### app.Router

This feature is not supported yet.

### app.Listen(port)

Bind and listen for connections on the given host and port, this method is [stackr.Listen()](http://gostackr.appspot.com/).

    app := f.CreateServer()
    app.Listen(3000)

If running in side a container such as the [Google App Engine](https://developers.google.com/appengine/), `app` can be passed to the standard [http.Handle()](http://golang.org/pkg/net/http/#Handle) function.

    http.Handle("/", app)

## Request

### req.Params

This feature is not supported yet.

### req.Query

This property is a map containing the first item of parsed query-string parameters.

    // GET /search?q=ric+allinson
    req.Query["q"]
    // => "ric allinson"

### req.Body

This property is a map containing the first item of the parsed request body. This feature is provided by the `http.PostForm` property, though other body parsing middleware may populate this property instead.

    // POST user=ric&email=ric@randomism.org
    req.Body["user"]
    // => "ric"

    req.Body["email"]
    // => "ric@randomism.org"

### req.Files

This feature is not supported yet.

### req.Param(name)

Return the value of param `name` when present.

    // ?name=ric
    req.Param("name")
    // => "ric"

    // POST name=ric
    req.Param("name")
    // => "ric"

Lookup is performed in the following order:

* req.Params
* req.Body
* req.Query

Direct access to `req.Body`, `req.Params`, and `req.Query` should be favored for clarity - unless you truly accept input from each object.

### req.Route

This feature is not supported yet.

### req.Cookie(name, [interface])

Returns the cookie value for `name` and optionally populates the given `interface`. Cookie values are URL and base64 encoded.

    // Cookie: foo=bar
    req.cookie("foo")
    // => "bar"

    // Cookie: foo=eyJmb28iOiJiYXIifQ%3D%3D
    var f map[string]interface{}
    t := req.Cookie("foo", &f)
    // f["foo"] == "bar"

### req.SignedCookie(name, [interface])

Contains the signed cookies sent by the user-agent, unsigned and ready for use. Signed cookies are accessed by a different function to show developer intent, otherwise a malicious attack could be placed on `req.Cookie` values which are easy to spoof. Note that signing a cookie does not mean it is "hidden" nor encrypted, this simply prevents tampering as the secret used to sign is private.

    // Cookie: foo=YmFyLld2WHdGQVBpaDNuQllfWUJhWWp3MmlONmN6VTFqam5MNjU1ZHZrcnFjbE09
    req.SignedCookie("foo")
    // => "bar"

    // Cookie: foo=eyJmb28iOiJiYXIifS5QU1hjUGdOS3NwZFR6Q3BmOW1qN2JFR2RTUUx3MU5nWTRkMkE2QXpFTktjPQ%3D%3D
    var f map[string]interface{}
    t := req.SignedCookie("foo", &f)
    // f["foo"] == "bar"

### req.Get(field)

Get the case-insensitive request header `field`.

    req.Get("Content-Type")
    // => "text/plain"

    req.Get("content-type")
    // => "text/plain"

    req.Get("Something")
    // => undefined

Alias for `req.Header.Get(field)`.

### req.Accepts(types)

Check if the given type is acceptable, returning true or false - in which case you should respond with 406 "Not Acceptable".

The type value must be a single mime type string such as "application/json".

    // Accept: text/html, application/json
    req.accepts("application/json")
    // => true

### req.Accepted()

Return an slice of Accepted media types ordered from highest quality to lowest.

    // Accept: "text/html, application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
    req.Accepted()
    // ["text/html", "application/xhtml+xml", "application/xml", "*.*"]

### req.Is(type)

Check if the incoming request contains the "Content-Type" header field, and it matches the give mime `type`.

    // With Content-Type: text/html; charset=utf-8
    req.Is("html")
    req.Is("text/html")
    // => true

    // When Content-Type is application/json
    req.Is("json")
    req.Is("application/json")
    // => true

    req.Is("html")
    // => false

### req.Ip

Return the remote address, or when "trust proxy" is enabled - the upstream address.

    req.Ip
    // => "127.0.0.1"

### req.Ips

When "trust proxy" is `true`, parse the "X-Forwarded-For" ip address list and return an array, otherwise an empty array is returned. For example if the value were "client, proxy1, proxy2" you would receive the array `["client", "proxy1", "proxy2"]` where "proxy2" is the furthest down-stream.

### req.Path

Returns the request URL pathname.

    // example.com/users?sort=desc
    req.Path
    // => "/users"

### req.Host

Returns the hostname from the "Host" header field.

    // Host: "example.com:3000"
    req.Host
    // => "example.com"

### req.Fresh()

Check if the request is fresh - aka Last-Modified and/or the ETag still match, indicating that the resource is "fresh".

    req.Fresh()
    // => true

### req.Stale()

Check if the request is stale - aka Last-Modified and/or the ETag do not match, indicating that the resource is "stale".

    req.Stale()
    // => true

### req.Xhr

Check if the request was issued with the "X-Requested-With" header field set to "XMLHttpRequest" (jQuery etc).

    req.Xhr
    // => true

### req.Protocol

Return the protocol string "http" or "https" when requested with TLS. When the "trust proxy" setting is enabled the "X-Forwarded-Proto" header field will be trusted. If you"re running behind a reverse proxy that supplies https for you this may be enabled.

    req.Protocol
    // => "http"

### req.Secure

Check if a TLS connection is established. This is a short-hand for:

    "https" == req.Protocol;

### req.Subdomains

This feature is not supported yet.

### req.OriginalUrl

This property is much like `req.Url`, however it retains the original request url, allowing you to rewrite `req.Url` freely for internal routing purposes. For example the "mounting" feature of `app.Use()` will rewrite `req.Url` to strip the mount point.

    // GET /search?q=something
    req.OriginalUrl
    // => "/search?q=something"

### req.AcceptedLanguages()

Return a slice of Accepted languages ordered from highest quality to lowest.

    Accept-Language: en;q=.5, en-us
    // => ["en-us", "en"]

### req.AcceptedCharsets()

Return a slice of Accepted charsets ordered from highest quality to lowest.

    Accept-Charset: iso-8859-5;q=.2, unicode-1-1;q=0.8
    // => ["unicode-1-1", "iso-8859-5"]

### req.AcceptsCharset(charset)

Check if the given charset is acceptable.

### req.AcceptsLanguage(lang)

Check if the given lang is acceptable.

## Response

### res.Status(code)

Alias of `stackr.StatusCode`.

    res.Status(404)

### res.Set(field, value)

Set header `field` to `value`.

res.set("Content-Type", "text/plain")

Alias of `http.ResponseWriter.Header().Set(field, value)`.

### res.Get(field)

Get the case-insensitive response header `field`.

    res.Get("Content-Type")
    // => "text/plain"

Alias of `http.ResponseWriter.Header().Get(field)`

### res.Cookie(name, value, [options])

Set cookie `name` to `value`, where `value` may be a string or an interface that will be converted to JSON. The path option defaults to "/". Options are passed as a (http.Cookie)[http://golang.org/pkg/net/http/#Cookie].

    res.Cookie("name", "ric")

The maxAge option is a convenience option for setting "expires" relative to the current time in milliseconds.

    res.Cookie("rememberme", "1", &http.Cookie{MaxAge: 900000, HttpOnly: true})

An interface may be passed which is then serialized as JSON.

    res.Cookie("cart", map[string]string{"name": "ric"})
    res.Cookie("cart", map[string]string{"name": "ric"}, &http.Cookie{MaxAge: 900000})

All cookie `values` are URL and base64 encoded.

### res.SignedCookie(name, value, [options])

Same as `res.Cookie(name, value, [options])` except that the cookie is signed.

### res.ClearCookie(name, [options])

Clear cookie `name`. The `path` option defaults to "/".

    res.Cookie("name", "ric", &http.Cookie{path: "/admin"})
    res.ClearCookie("name", &http.Cookie{path: "/admin"})

### res.Redirect(url, [status])

Redirect to the given `url` with optional `status` code defaulting to 302 "Found".

    res.Redirect("/foo/bar")
    res.Redirect("http://example.com")
    res.Redirect("http://example.com", 301)
    res.Redirect("../login")

Forgery supports a few forms of redirection, first being a fully qualified URI for redirecting to a different site:

    res.Redirect("http://yahoo.com")

The second form is the pathname-relative redirect, for example if you were on `http://example.com/admin/post/new`, the following redirect to `/admin` would land you at `http://example.com/admin`:

    res.Redirect("/admin")

This next redirect is relative to the `mount` point of the application. For example if you have a blog application mounted at `/blog`, ideally it has no knowledge of where it was mounted, so where a redirect of /admin/post/new would simply give you `http://example.com/admin/post/new`, the following mount-relative redirect would give you `http://example.com/blog/admin/post/new`:

    res.Redirect("admin/post/new")

Pathname relative redirects are also possible. If you were on `http://example.com/admin/post/new`, the following redirect would land you at `http//example.com/admin/post`:

    res.Redirect("..")

The final special-case is a back redirect, redirecting back to the Referer, defaulting to / when missing.

    res.Redirect("back")

### res.Location(uri)

Set the location header.

    res.Location("/foo/bar")
    res.Location("foo/bar")
    res.Location("http://example.com")
    res.Location("../login")
    res.Location("back")

You can use the same kind of urls as in res.Redirect().

### res.Charset

Assign the `charset`. Defaults to "utf-8".

    res.charset = "value";
    res.send("some html")
    // => Content-Type: text/html; charset=value

### res.Send()

Send a response.

    res.Send([]byte{114, 105, 99}])
    res.Send(map[string]string{"some": "json"})
    res.Send("some html")
    res.Send("Sorry, we cannot find that!", 404)
    res.Send(map[string]string{"error": "msg"}, 500)
    res.Send(200)

This method performs a myriad of useful tasks for simple non-streaming responses such as automatically assigning the Content-Length unless previously defined and providing automatic __HEAD__ and HTTP cache freshness support.

When a `string` or `[]byte` is given the Content-Type is set defaulted to "text/html":

    res.Send("some html")

When an `interface` is given forgery will respond with the JSON representation:

    res.Send(map[string]string{"user": "ric"})
    res.Send([]int{1, 2, 3})

Finally when a `int` is given without any of the previously mentioned bodies, then a response body string is assigned for you. For example 200 will respond will the text "OK", and 404 "Not Found" and so on.

    res.Send(200)
    res.Send(404)
    res.Send(500)

### res.Json(body, [status])

Send a JSON response. This method is identical to `res.Send()` when an `interface` is passed, however it may be used for explicit JSON conversion of non-objects (null, undefined, etc), though these are technically not valid JSON.

    res.Json(null)
    res.Json(map[string]string{"user": "ric"})
    res.Json(map[string]string{"error": "message"}, 500)

### res.Jsonp(body, [status])

Send a JSON response with JSONP support. This method is identical to res.json() however opts-in to JSONP callback support.

    res.jsonp(null)
    // => null

    res.Json(map[string]string{"user": "ric"})
    // => {"user": "ric"}

    res.Json(map[string]string{"error": "message"}, 500)
    // => {"error": "message"}

By default the JSONP callback name is simply `callback`, however you may alter this with the _jsonp callback name_ setting. The following are some examples of JSONP responses using the same code:

    // ?callback=foo
    res.Json(map[string]string{"user": "ric"})
    // => foo({"user": "ric"})

    app.set("jsonp callback name", "cb")

    // ?cb=foo
    res.Json(map[string]string{"error": "message"}, 500)
    // => foo({ "error": "message" })

### res.ContentType(type)

Sets the Content-Type to the mime lookup of `type`, or when "/" is present the Content-Type is simply set to this literal value.

    res.ContentType(".html")
    res.ContentType("html")
    res.ContentType("json")
    res.ContentType("application/json")
    res.ContentType("png")

### res.Format(map[string]func())

Performs content-negotiation on the request Accept header field when present. This method uses `req.Accepted()`, a slice of acceptable types ordered by their quality values, otherwise the first callback is invoked. When no match is performed the server responds with 406 "Not Acceptable", or invokes the `default` callback.

The Content-Type is set for you when a callback is selected, however you may alter this within the callback using `res.Set()` or `res.ContentType()`.

The following example would respond with `{"message": "hey"}` when the Accept header field is set to "application/json" or "*/json", however if "*/*" is given then "hey" will be the response.

    res.Format(map[string]func() {
        "text/plain": func() {
            res.Send("hey")
        },

        "text/html": func() {
            res.Send("hey")
        },

        "application/json": func() {
            res.Send(map[string]string{"message": "hey"})
        }
    })

### res.Attachment([filename])

Sets the Content-Disposition header field to "attachment". If a `filename` is given then the Content-Type will be automatically set based on the extname via `res.ContentType()`, and the Content-Disposition"s "filename=" parameter will be set.

    res.Attachment()
    // Content-Disposition: attachment

    res.Attachment("path/to/logo.png")
    // Content-Disposition: attachment; filename="logo.png"
    // Content-Type: image/png

### res.Sendfile(path)

Transfer the file at the given `path`. Alias for [http.ServeFile](http://golang.org/pkg/net/http/#ServeFile).

### res.Download(path, [filename])

Transfer the file at `path` as an "attachment", typically browsers will prompt the user for download. The Content-Disposition "filename=" parameter, aka the one that will appear in the browser dialog is set to `path` by default, however you may provide an override `filename`.

    res.Download("/report-12345.pdf")

    res.Download("/report-12345.pdf", "report.pdf")

Uses `req.Send()` to do the file transfer.

### res.Links(link, rel)

Join the given `link`, `rel` to populate the "Link" response header field.

    res.Links("http://api.example.com/users?page=2", "next")
    res.Links("http://api.example.com/users?page=5", "last")

yields:

    Link: <http://api.example.com/users?page=2>; rel="next", 
          <http://api.example.com/users?page=5>; rel="last"

### res.Locals

Response local variables are scoped to the request, thus only available to the view(s) rendered during that request / response cycle, if any. Otherwise this API is identical to `app.Locals`.

This object is useful for exposing request-level information such as the request pathname, authenticated user, user settings etcetera.

    app.All(func(req *f.Request, res *f.Response, next func()) {
        res.Locals["user"] = req.Map["user"]
        res.Locals["authenticated"] = req.Map["authenticated"]
        next()
    })

### res.Render(view, [locals...])

Render the `view` file responding with the rendered string using `res.Send()`. The `view` file is located using the `views` setting.

    res.Render("index.html", map[string]string{
        "body": "Document body",
    })

For the example below forgery will look for the file `./views/page.html` and attempt to render it using the `view engine` assigned to the extension `.html`.

    res.Render("page.html", map[string]string{
        "body": "Document body",
    }, map[string]int{
        "prev": 5,
        "next": 11,
    })
