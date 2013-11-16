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

    app.Set("title", "My Site");
    app.Get("title");
    // => "My Site"

### app.Get(name)

Get setting `name` value.

    app.Get("title");
    // => ""

    app.Set("title", "My Site");
    app.Get("title");
    // => "My Site"

### app.Enable(name)

Set setting `name` to `true`.

    app.Enable("trust proxy");
    app.Get("trust proxy");
    // => "true"

### app.Disable(name)

Set setting `name` to `false`.

    app.Disable("trust proxy");
    app.Get("trust proxy");
    // => "false"

### app.Enabled(name)

Check if setting `name` is enabled.

    app.Enabled("trust proxy");
    // => false

    app.Enable("trust proxy");
    app.Enabled("trust proxy");
    // => true

### app.Disabled(name)

Check if setting `name` is disabled.

    app.Disabled("trust proxy");
    // => true

    app.Enable("trust proxy");
    app.Disabled("trust proxy");
    // => false

### app.Configure([env], callback)

Conditionally invoke `callback` when `env` matches `app.get("env")`. This method is effectively an `if` statement as illustrated in the following snippets. These functions are __not__ required in order to use `app.set()` and other configuration methods.

    // all environments
    app.Configure(func() {
        app.Set("title", "My Application");
    })

    // development only
    app.Configure("development", func(){
        app.Set("db uri", "localhost/dev");
    })

    // production only
    app.Configure("production", func(){
        app.Set("db uri", "n.n.n.n/prod");
    })

effectively sugar for:

    // all environments
    app.Set("title", "My Application");

    // development only
    if "development" == app.Get("env") {
        app.Set("db uri", "localhost/dev");
    }

    // production only
    if "production" == app.Get("env") {
        app.Set("db uri", "n.n.n.n/prod");
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

    app.all("*", requireAuthentication, loadUser);

Or the equivalent:

    app.all("*", requireAuthentication)
    app.all("*", loadUser);

Another great example of this is white-listed "global" functionality. Here the example is much like before, however only restricting paths prefixed with "/api":

    app.all("/api/", requireAuthentication);


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
    req.Is("html");
    req.Is("text/html");
    // => true

    // When Content-Type is application/json
    req.Is("json");
    req.Is("application/json");
    // => true

    req.Is("html");
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

* res.Charset
* res.Locals
* res.Status()
* res.ContentType()
* res.Get()
* res.Set()
* res.Send()
* res.Json()
* res.Jsonp()
* res.Render()
* res.Redirect()
* res.Cookie()
* req.SignedCookie()
* res.ClearCookie()
* res.Sendfile()
* res.Download()
* res.Format()
* res.Links()
* res.Vary()
* res.Attachment()
* res.Location()
