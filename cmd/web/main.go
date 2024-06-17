package main

/*
	Responsibilities of our main() function are limited to:

	Parsing the runtime configuration settings for the application;
	Establishing the dependencies for the handlers; and
	Running the HTTP server.
*/

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

/*
	how can we make our new structured logger available to our
	home function from main()?

	And this question generalizes further. Most web applications will have multiple
	dependencies that their handlers need to access, such as a database connection pool,
	centralized error handlers, and template caches. What we really want to answer is: how can
	we make any dependency available to our handlers?
	There are a few different ways to do this, the simplest being to just put the dependencies in
	global variables. But in general, it is good practice to inject dependencies into your handlers.
	It makes your code more explicit, less error-prone, and easier to unit test than if you use
	global variables.

	For applications where all your handlers are in the same package, like ours, a neat way to
	inject dependencies is to put them into a custom application struct, and then define your
	handler functions as methods against application.
*/
// Define an application struct to hold the application-wide dependencies for the
// web applicion.
type application struct {
	logger *slog.Logger
}

func main() {

	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// Use the slog.New() function to initialize a new structured logger, which
	// writes to the standard out stream and uses the default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	// Initialize a new instance of our application struct, containing the
	// dependencies (for now, just the structured logger)
	app := &application{
		logger: logger,
	}

	// Print a log message to say that the server is starting.
	logger.Info("starting server on", "addr", cfg.addr)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux that we get from app.routes().
	// If http.ListenAndServe() returns an error
	// we use the logger.Error() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
