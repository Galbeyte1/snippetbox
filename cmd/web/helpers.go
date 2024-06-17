package main

import (
	"net/http"
)

// tHE SERVEReRROR HELPER WRITES A LOG ENTRY AT error level (including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		// Use debug.Stack() to get the stack trace outlining the execution path of
		// the application for the current goroutine. This returns a byte slice, which
		// we need to convert to a string so that it's readable in the log entry.
		// trace = string(debug.Stack()) // make sure to add it to app.logger.Error
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(SERVER_ERROR), SERVER_ERROR)
}

// The clientError helper sends a specific status code and corresponding description
// to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
