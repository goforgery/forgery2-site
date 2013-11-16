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

    app.all('*', requireAuthentication, loadUser);

Or the equivalent:

    app.all('*', requireAuthentication)
    app.all('*', loadUser);

Another great example of this is white-listed "global" functionality. Here the example is much like before, however only restricting paths prefixed with "/api":

    app.all('/api/', requireAuthentication);


### app.Locals

Application local variables are provided to all templates rendered within the application. This is useful for providing app-level data.

    app.Locals["title'] = "My App";

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

* req.Ip
* req.Ips
* req.Path
* req.Xhr
* req.Protocol
* req.Secure
* req.Params
* __TBD__ req.Route
* req.Get()
* req.Param()
* req.Cookie()
* req.SignedCookie()
* req.Fresh()
* req.Stale()
* req.Is()
* req.Accepts()
* req.AcceptsCharset()
* req.AcceptsLanguage()
* req.Accepted()
* req.AcceptedLanguages()
* req.AcceptedCharsets()

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
