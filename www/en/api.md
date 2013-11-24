# API Reference

* __[f.CreateServer](#f.CreateServer)__
* __[Application](#Application)__
* [app.Set()](#app.Set)
* [app.Get()](#app.Get)
* [app.Enable()](#app.Enable)
* [app.Disable()](#app.Disable)
* [app.Configure()](#app.Configure)
* [Application Settings](#Application_Settings)
* [app.Use()](#app.Use)
* [app.Engine()](#app.Engine)
* [app.Param()](#app.Param)
* [Application Routing](#app.VERB)
* [app.All()](#app.All)
* [app.Locals](#app.Locals)
* [app.Render()](#app.Render)
* [app.Listen()](#app.Listen)
* __[f.Request](#f.Request)__
* [req.Params](#req.Params)
* [req.Query](#req.Query)
* [req.Body](#req.Body)
* [req.Files](#req.Files)
* [req.Param()](#req.Param)
* [req.Route](#req.Route)
* [req.Cookie()](#req.Cookie)
* [req.SignedCookie()](#req.SignedCookie)
* [req.Get()](#req.Get)
* [req.Accepts()](#req.Accepts)
* [req.Accepted()](#req.Accepted)
* [req.Is()](#req.Is)
* [req.Ip](#req.Ip)
* [req.Ips](#req.Ips)
* [req.Path](#req.Path)
* [req.Host](#req.Host)
* [req.Fresh()](#req.Fresh)
* [req.Stale()](#req.Stale)
* [req.Xhr](#req.Xhr)
* [req.Protocol](#req.Protocol)
* [req.Secure](#req.Secure)
* [req.Subdomains()](#req.Subdomains)
* [req.OriginalUrl](#req.OriginalUrl)
* [req.AcceptedLanguages()](#req.AcceptedLanguages)
* [req.AcceptedCharsets()](#req.AcceptedCharsets)
* [req.AcceptsCharset()](#req.AcceptsCharset)
* [req.AcceptsLanguage()](#req.AcceptsLanguage)
* __[f.Response](#f.Response)__
* [res.Status()](#res.Status)
* [res.Set()](#res.Set)
* [res.Get()](#res.Get)
* [res.Cookie()](#res.Cookie)
* [res.SignedCookie()](#res.SignedCookie)
* [res.ClearCookie()](#res.ClearCookie)
* [res.Redirect()](#res.Redirect)
* [res.Location()](#res.Location)
* [res.Charset](#res.Charset)
* [res.Send()](#res.Send)
* [res.Json()](#res.Json)
* [res.Jsonp()](#res.Jsonp)
* [res.ContentType()](#res.ContentType)
* [res.Format()](#res.Format)
* [res.Attachment()](#res.Attachment)
* [res.Sendfile()](#res.Sendfile)
* [res.Download()](#res.Download)
* [res.Links()](#res.Links)
* [res.Locals](#res.Locals)
* [res.Render()](#res.Render)
* __[Middleware](#Middleware)__
* [f.ErrorHandler()](#f.ErrorHandler)
* [f.Favicon()](#f.Favicon)
* [f.Logger()](#f.Logger)
* [f.MethodOverride()](#f.MethodOverride)
* [f.ResponseTime()](#f.ResponseTime)
* [f.Static()](#f.Static)

## <a class="jump" name="f.CreateServer"></a>f.CreateServer()

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

## <a class="jump" name="Application"></a>Application

### <a class="jump" name="app.Set"></a>app.Set(name, value)

Assigns setting `name` to `value`.

    app.Set("title", "My Site")
    app.Get("title")
    // => "My Site"

### <a class="jump" name="app.Get"></a>app.Get(name)

Get setting `name` value.

    app.Get("title")
    // => ""

    app.Set("title", "My Site")
    app.Get("title")
    // => "My Site"

### <a class="jump" name="app.Enable"></a>app.Enable(name)

Set setting `name` to `true`.

    app.Enable("trust proxy")
    app.Get("trust proxy")
    // => "true"

### <a class="jump" name="app.Disable"></a>app.Disable(name)

Set setting `name` to `false`.

    app.Disable("trust proxy")
    app.Get("trust proxy")
    // => "false"

### <a class="jump" name="app.Enabled"></a>app.Enabled(name)

Check if setting `name` is enabled.

    app.Enabled("trust proxy")
    // => false

    app.Enable("trust proxy")
    app.Enabled("trust proxy")
    // => true

### <a class="jump" name="app.Disabled"></a>app.Disabled(name)

Check if setting `name` is disabled.

    app.Disabled("trust proxy")
    // => true

    app.Enable("trust proxy")
    app.Disabled("trust proxy")
    // => false

### <a class="jump" name="app.Configure"></a>app.Configure([env], callback)

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

### <a class="jump" name="Application_Settings"></a>Application Settings

The following settings are provided to alter how Forgery will behave:

* _env_ Environment mode, defaults to `os.Getenv("GO_ENV")` or "development"
* _trust proxy_ Enables reverse proxy support, disabled by default
* _jsonp callback name_ Changes the default callback name of _?callback=_
* _json spaces_ JSON response spaces for formatting, defaults to _"  "_ (2 spaces) in development, _0_ in production
* TBD _case sensitive routing_ Enable case sensitivity, disabled by default, treating "/Foo" and "/foo" as the same
* TBD _strict routing_ Enable strict routing, by default "/foo" and "/foo/" are treated the same by the router
* TBD _view cache_ Enables view template compilation caching, enabled in production by default
* _view engine_ The default engine extension to use when omitted
* _views_ The view directory path, defaulting to "./views"

### <a class="jump" name="app.Use"></a>app.Use([path], function)

Wrapper for `stackr.Use()`. See the [Stackr](http://gostackr.appspot.com/) documentation for details.

### <a class="jump" name="app.Engine"></a>app.Engine(ext, callback)

Register the given template engine `callback` as `ext`.

    app.Engine(".html", fmustache.Make())

### <a class="jump" name="app.Param"></a>TBD app.Param([name], callback)

This feature is not supported yet.

### <a class="jump" name="app.VERB"></a>app.VERB(path, [callback...], callback)

The `app.VERB()` methods provide the routing functionality in Forgery, where __VERB__ is one of the HTTP verbs, such as `app.Post()`. Multiple callbacks may be given, all are treated equally, and behave just like middleware, with the one exception that these callbacks may invoke `next()` to bypass the remaining route callback(s). This mechanism can be used to perform pre-conditions on a route then pass control to subsequent routes when there is no reason to proceed with the route matched.

The following snippet illustrates the most simple route definition possible. Forgery translates the path strings to regular expressions, used internally to match incoming requests. Query strings are not considered when performing these matches, for example "GET /" would match the following route, as would "GET /?name=ric".

    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Send("Hello world.")
    })

Several callbacks may also be passed, useful for performing validations, loading data, etc.

    app.Get("/user/:id", loadUser, func(req *f.Request, res *f.Response, next func()) {
        // ...
    })

__NOTE: Regular expressions not supported yet__.

### <a class="jump" name="app.All"></a>app.All(path, [callback...], callback)

This method functions just like the `app.VERB()` methods, however it matches all HTTP verbs.

This method is extremely useful for mapping "global" logic for specific path prefixes or arbitrary matches. For example if you placed the following route at the top of all other route definitions, it would require that all routes from that point on would require authentication, and automatically load a user. Keep in mind that these callbacks do not have to act as end points, `loadUser` can perform a task, then `next()` to continue matching subsequent routes.

    app.all("*", requireAuthentication, loadUser)

Or the equivalent:

    app.all("*", requireAuthentication)
    app.all("*", loadUser)

Another great example of this is white-listed "global" functionality. Here the example is much like before, however only restricting paths prefixed with "/api":

    app.all("/api/", requireAuthentication)


### <a class="jump" name="app.Locals"></a>app.Locals

Application local variables are provided to all templates rendered within the application. This is useful for providing app-level data.

    app.Locals["title"] = "My App";

### <a class="jump" name="app.Render"></a>app.Render(view, [locals...])

Render a `view` file returning with the rendered string. This is the app-level variant of `res.Render()`, and otherwise behaves the same way.

    s := res.Render("index.html", map[string]string{
        "body": "Document body",
    })

### <a class="jump" name="app.Router"></a>app.Router

This feature is not supported yet.

### <a class="jump" name="app.Listen"></a>app.Listen(port)

Bind and listen for connections on the given host and port, this method is [stackr.Listen()](http://gostackr.appspot.com/).

    app := f.CreateServer()
    app.Listen(3000)

If running in side a container such as the [Google App Engine](https://developers.google.com/appengine/), `app` can be passed to the standard [http.Handle()](http://golang.org/pkg/net/http/#Handle) function.

    http.Handle("/", app)

## <a class="jump" name="f.Request"></a>f.Request

### <a class="jump" name="req.Params"></a>req.Params

This property is map containing properties mapped to the named route "parameters". For example if you have the route `/user/:name`, then the "name" property is available to you as `req.params["name"]`.

    // GET /user/ric
    req.params["name"]
    // => "ric"

### <a class="jump" name="req.Query"></a>req.Query

This property is a map containing the first item of parsed query-string parameters.

    // GET /search?q=ric+allinson
    req.Query["q"]
    // => "ric allinson"

### <a class="jump" name="req.Body"></a>req.Body

This property is a map containing the first item of the parsed request body. This feature is provided by the `http.PostForm` property, though other body parsing middleware may populate this property instead.

    // POST user=ric&email=ric@randomism.org
    req.Body["user"]
    // => "ric"

    req.Body["email"]
    // => "ric@randomism.org"

### <a class="jump" name="req.Files"></a>req.Files

This feature is not supported yet.

### <a class="jump" name="req.Param"></a>req.Param(name)

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

### <a class="jump" name="req.Route"></a>req.Route

This feature is not supported yet.

### <a class="jump" name="req.Cookie"></a>req.Cookie(name, [interface])

Returns the cookie value for `name` and optionally populates the given `interface`. Cookie values are URL and base64 encoded.

    // Cookie: foo=bar
    req.cookie("foo")
    // => "bar"

    // Cookie: foo=eyJmb28iOiJiYXIifQ%3D%3D
    var f map[string]interface{}
    t := req.Cookie("foo", &f)
    // f["foo"] == "bar"

### <a class="jump" name="req.SignedCookie"></a>req.SignedCookie(name, [interface])

Contains the signed cookies sent by the user-agent, unsigned and ready for use. Signed cookies are accessed by a different function to show developer intent, otherwise a malicious attack could be placed on `req.Cookie` values which are easy to spoof. Note that signing a cookie does not mean it is "hidden" nor encrypted, this simply prevents tampering as the secret used to sign is private.

    // Cookie: foo=YmFyLld2WHdGQVBpaDNuQllfWUJhWWp3MmlONmN6VTFqam5MNjU1ZHZrcnFjbE09
    req.SignedCookie("foo")
    // => "bar"

    // Cookie: foo=eyJmb28iOiJiYXIifS5QU1hjUGdOS3NwZFR6Q3BmOW1qN2JFR2RTUUx3MU5nWTRkMkE2QXpFTktjPQ%3D%3D
    var f map[string]interface{}
    t := req.SignedCookie("foo", &f)
    // f["foo"] == "bar"

### <a class="jump" name="req.Get"></a>req.Get(field)

Get the case-insensitive request header `field`.

    req.Get("Content-Type")
    // => "text/plain"

    req.Get("content-type")
    // => "text/plain"

    req.Get("Something")
    // => undefined

Alias for `req.Header.Get(field)`.

### <a class="jump" name="req.Accepts"></a>req.Accepts(types)

Check if the given type is acceptable, returning true or false - in which case you should respond with 406 "Not Acceptable".

The type value must be a single mime type string such as "application/json".

    // Accept: text/html, application/json
    req.accepts("application/json")
    // => true

### <a class="jump" name="req.Accepted"></a>req.Accepted()

Return an slice of Accepted media types ordered from highest quality to lowest.

    // Accept: "text/html, application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
    req.Accepted()
    // ["text/html", "application/xhtml+xml", "application/xml", "*.*"]

### <a class="jump" name="req.Is"></a>req.Is(type)

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

### <a class="jump" name="req.Ip"></a>req.Ip

Return the remote address, or when "trust proxy" is enabled - the upstream address.

    req.Ip
    // => "127.0.0.1"

### <a class="jump" name="req.Ips"></a>req.Ips

When "trust proxy" is `true`, parse the "X-Forwarded-For" ip address list and return an array, otherwise an empty array is returned. For example if the value were "client, proxy1, proxy2" you would receive the array `["client", "proxy1", "proxy2"]` where "proxy2" is the furthest down-stream.

### <a class="jump" name="req.Path"></a>req.Path

Returns the request URL pathname.

    // example.com/users?sort=desc
    req.Path
    // => "/users"

### <a class="jump" name="req.Host"></a>req.Host

Returns the hostname from the "Host" header field.

    // Host: "example.com:3000"
    req.Host
    // => "example.com"

### <a class="jump" name="req.Fresh"></a>req.Fresh()

Check if the request is fresh - aka Last-Modified and/or the ETag still match, indicating that the resource is "fresh".

    req.Fresh()
    // => true

### <a class="jump" name="req.Stale"></a>req.Stale()

Check if the request is stale - aka Last-Modified and/or the ETag do not match, indicating that the resource is "stale".

    req.Stale()
    // => true

### <a class="jump" name="req.Xhr"></a>req.Xhr

Check if the request was issued with the "X-Requested-With" header field set to "XMLHttpRequest" (jQuery etc).

    req.Xhr
    // => true

### <a class="jump" name="req.Protocol"></a>req.Protocol

Return the protocol string "http" or "https" when requested with TLS. When the "trust proxy" setting is enabled the "X-Forwarded-Proto" header field will be trusted. If you"re running behind a reverse proxy that supplies https for you this may be enabled.

    req.Protocol
    // => "http"

### <a class="jump" name="req.Secure"></a>req.Secure

Check if a TLS connection is established. This is a short-hand for:

    "https" == req.Protocol;

### <a class="jump" name="req.Subdomains"></a>req.Subdomains()

Return subdomain as a slice of strings. Subdomains are the dot-separated parts of the host before the main domain of the app. By default, the domain of the app is assumed to be the last two parts of the host. This can be changed by setting "subdomain offset".

    // Host: ric.allinson.example.com
    req.Subdomain()
    // => ["allinson", "ric"]

### <a class="jump" name="req.OriginalUrl"></a>req.OriginalUrl

This property is much like `req.Url`, however it retains the original request url, allowing you to rewrite `req.Url` freely for internal routing purposes. For example the "mounting" feature of `app.Use()` will rewrite `req.Url` to strip the mount point.

    // GET /search?q=something
    req.OriginalUrl
    // => "/search?q=something"

### <a class="jump" name="req.AcceptedLanguages"></a>req.AcceptedLanguages()

Return a slice of Accepted languages ordered from highest quality to lowest.

    Accept-Language: en;q=.5, en-us
    // => ["en-us", "en"]

### <a class="jump" name="req.AcceptedCharsets"></a>req.AcceptedCharsets()

Return a slice of Accepted charsets ordered from highest quality to lowest.

    Accept-Charset: iso-8859-5;q=.2, unicode-1-1;q=0.8
    // => ["unicode-1-1", "iso-8859-5"]

### <a class="jump" name="req.AcceptsCharset"></a>req.AcceptsCharset(charset)

Check if the given charset is acceptable.

### <a class="jump" name="req.AcceptsLanguage"></a>req.AcceptsLanguage(lang)

Check if the given lang is acceptable.

## <a class="jump" name="f.Response"></a>f.Response

### <a class="jump" name="res.Status"></a>res.Status(code)

Alias of `stackr.StatusCode`.

    res.Status(404)

### <a class="jump" name="res.Set"></a>res.Set(field, value)

Set header `field` to `value`.

res.set("Content-Type", "text/plain")

Alias of `http.ResponseWriter.Header().Set(field, value)`.

### <a class="jump" name="res.Get"></a>res.Get(field)

Get the case-insensitive response header `field`.

    res.Get("Content-Type")
    // => "text/plain"

Alias of `http.ResponseWriter.Header().Get(field)`

### <a class="jump" name="res.Cookie"></a>res.Cookie(name, value, [options])

Set cookie `name` to `value`, where `value` may be a string or an interface that will be converted to JSON. The path option defaults to "/". Options are passed as a (http.Cookie)[http://golang.org/pkg/net/http/#Cookie].

    res.Cookie("name", "ric")

The maxAge option is a convenience option for setting "expires" relative to the current time in milliseconds.

    res.Cookie("rememberme", "1", &http.Cookie{MaxAge: 900000, HttpOnly: true})

An interface may be passed which is then serialized as JSON.

    res.Cookie("cart", map[string]string{"name": "ric"})
    res.Cookie("cart", map[string]string{"name": "ric"}, &http.Cookie{MaxAge: 900000})

All cookie `values` are URL and base64 encoded.

### <a class="jump" name="res.SignedCookie"></a>res.SignedCookie(name, value, [options])

Same as `res.Cookie(name, value, [options])` except that the cookie is signed.

### <a class="jump" name="res.ClearCookie"></a>res.ClearCookie(name, [options])

Clear cookie `name`. The `path` option defaults to "/".

    res.Cookie("name", "ric", &http.Cookie{path: "/admin"})
    res.ClearCookie("name", &http.Cookie{path: "/admin"})

### <a class="jump" name="res.Redirect"></a>res.Redirect(url, [status])

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

### <a class="jump" name="res.Location"></a>res.Location(uri)

Set the location header.

    res.Location("/foo/bar")
    res.Location("foo/bar")
    res.Location("http://example.com")
    res.Location("../login")
    res.Location("back")

You can use the same kind of urls as in res.Redirect().

### <a class="jump" name="res.Charset"></a>res.Charset

Assign the `charset`. Defaults to "utf-8".

    res.charset = "value";
    res.send("some html")
    // => Content-Type: text/html; charset=value

### <a class="jump" name="res.Send"></a>res.Send()

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

### <a class="jump" name="res.Json"></a>res.Json(body, [status])

Send a JSON response. This method is identical to `res.Send()` when an `interface` is passed, however it may be used for explicit JSON conversion of non-objects (null, undefined, etc), though these are technically not valid JSON.

    res.Json(null)
    res.Json(map[string]string{"user": "ric"})
    res.Json(map[string]string{"error": "message"}, 500)

### <a class="jump" name="res.Jsonp"></a>res.Jsonp(body, [status])

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

### <a class="jump" name="res.ContentType"></a>res.ContentType(type)

Sets the Content-Type to the mime lookup of `type`, or when "/" is present the Content-Type is simply set to this literal value.

    res.ContentType(".html")
    res.ContentType("html")
    res.ContentType("json")
    res.ContentType("application/json")
    res.ContentType("png")

### <a class="jump" name="res.Format"></a>res.Format(map[string]func())

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

### <a class="jump" name="res.Attachment"></a>res.Attachment([filename])

Sets the Content-Disposition header field to "attachment". If a `filename` is given then the Content-Type will be automatically set based on the extname via `res.ContentType()`, and the Content-Disposition"s "filename=" parameter will be set.

    res.Attachment()
    // Content-Disposition: attachment

    res.Attachment("path/to/logo.png")
    // Content-Disposition: attachment; filename="logo.png"
    // Content-Type: image/png

### <a class="jump" name="res.Sendfile"></a>res.Sendfile(path)

Transfer the file at the given `path`. Alias for [http.ServeFile](http://golang.org/pkg/net/http/#ServeFile).

### <a class="jump" name="res.Download"></a>res.Download(path, [filename])

Transfer the file at `path` as an "attachment", typically browsers will prompt the user for download. The Content-Disposition "filename=" parameter, aka the one that will appear in the browser dialog is set to `path` by default, however you may provide an override `filename`.

    res.Download("/report-12345.pdf")

    res.Download("/report-12345.pdf", "report.pdf")

Uses `req.Send()` to do the file transfer.

### <a class="jump" name="res.Links"></a>res.Links(link, rel)

Join the given `link`, `rel` to populate the "Link" response header field.

    res.Links("http://api.example.com/users?page=2", "next")
    res.Links("http://api.example.com/users?page=5", "last")

yields:

    Link: <http://api.example.com/users?page=2>; rel="next", 
          <http://api.example.com/users?page=5>; rel="last"

### <a class="jump" name="res.Locals"></a>res.Locals

Response local variables are scoped to the request, thus only available to the view(s) rendered during that request / response cycle, if any. Otherwise this API is identical to `app.Locals`.

This object is useful for exposing request-level information such as the request pathname, authenticated user, user settings etcetera.

    app.All(func(req *f.Request, res *f.Response, next func()) {
        res.Locals["user"] = req.Map["user"]
        res.Locals["authenticated"] = req.Map["authenticated"]
        next()
    })

### <a class="jump" name="res.Render"></a>res.Render(view, [locals...])

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

## <a class="jump" name="Middleware"></a>Middleware

### <a class="jump" name="f.ErrorHandler"></a>f.ErrorHandler()

Convenience attribute for accessing [stackr.ErrorHandler](http://godoc.org/github.com/ricallinson/stackr#ErrorHandler).

### <a class="jump" name="f.Favicon"></a>f.Favicon()

Convenience attribute for accessing [stackr.Favicon](http://godoc.org/github.com/ricallinson/stackr#Favicon).

### <a class="jump" name="f.Logger"></a>f.Logger()

Convenience attribute for accessing [stackr.Logger](http://godoc.org/github.com/ricallinson/stackr#Logger).

### <a class="jump" name="f.MethodOverride"></a>f.MethodOverride()

Convenience attribute for accessing [stackr.MethodOverride](http://godoc.org/github.com/ricallinson/stackr#MethodOverride).

### <a class="jump" name="f.ResponseTime"></a>f.ResponseTime()

Convenience attribute for accessing [stackr.ResponseTime](http://godoc.org/github.com/ricallinson/stackr#ResponseTime).

### <a class="jump" name="f.Static"></a>f.Static()

Convenience attribute for accessing [stackr.Static](http://godoc.org/github.com/ricallinson/stackr#Static).
