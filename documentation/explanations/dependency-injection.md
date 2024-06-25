# Structured Logging, Dependency Injection, and Centralized Error Handling

---

### Why Structured Logging ?

I was previously outputting log entries from my code using the `log.Printf()` and
`log.Fatal()` functions but now I started to use Go's `log/slog` package from it's standard library which let's me create custom structured loggers.

For many applications, using the standard logger will be good enough, and there’s no need to do anything more complex. But for applications which do a lot of logging, I want to make the log entries easier to filter and work with.

This allows for

- Distinguishing between different _severities_ of log entries
- Enforce a consistent structure for log entries so it's easier to parse when using external programs like Datadog, New Relic etc.

I'll use the `slog.New()` function to initialize a new structured logger, which writes to the standard out stream and uses the default settings. It's important to note custom loggers created by `slog.New()` are **concurrency-safe**. We can share a single logger and use it across multiple goroutines and in our HTTP handlers without needing to worry about race conditions.

```Go
logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level:     slog.LevelDebug,
    AddSource: true,
}))
```

I can then write a log entry at a specific severity
level by calling the `Debug()`, `Info()`, `Warn()` or `Error()`

---

### Dependency Injection

How can we make our new structured logger available to our `func home(w http.ResponseWriter, r *http.Request)` function from `func main()`?

How can we make **ANY** dependency available to our handlers?
There are a few different ways to do this, the simplest being to just make the dependency as a global variable. But in general, it is good practice to inject dependencies into your handlers. It makes your code more

- explicit
- less error-prone
- easier to unit test

than if you use global variables.

Most web applications will have multiple dependencies that their handlers need to access, such as a

- database connection pool,
- centralized error handlers,
- and template caches.

For applications where all the handlers are in the same package, like mine, a neat way to inject dependencies is to put them into a custom application struct, and then define the handler functions as methods against application.

I defined an application struct to hold the application-wide dependencies for the
web application.

```Go
    type application struct {
        logger   *slog.Logger
        snippets *models.SnippetModel
    }
```

I initialized a new instance of our application struct, containing the dependencies

```Go
    app := &application{
        logger:   logger,
        snippets: &models.SnippetModel{DB: db},
    }
```

```Go
func (app *application) home(w http.ResponseWriter, r *http.Request) {
```

Because the home handler is now a method against the application struct it can access its fields, including the structured logger. I can use this to create a log entry at Error level containing the error message, also including the request method and URI as attributes to assist with debugging. Let's introduce better error first.

### Centralized Error Handling

To seperate concerns and stop repeating code as I progress through the build processes I've dedicated error handling to some helper functions

The `serverError` is a method against the application struct and it's a helper which writes a log entries at Error level.

```Go
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
```

It writes a log entry at this level including the request method and URI as attributes, then sends a generic 500 Internal Server Error response to the user.

```Go
...
    var (
        method = r.Method
        uri = r.URL.RequestURI()

    )

    app.logger.Error(err.Error(), "method", method, "uri", uri)
    http.Error(w, http.StatusText(SERVER_ERROR), SERVER_ERROR)
```

Now as an example, within our `func (app *application) home` handler, we can pass simple centralized, structured error logging as such.

```Go
    ...
    ts, err := template.ParseFiles(files...)
	if err != nil {
        app.serverError(w, r, err)
    }
    ...
```
