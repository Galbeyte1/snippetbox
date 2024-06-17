package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const (
	NOT_ALLOWED  = http.StatusMethodNotAllowed
	OK           = http.StatusOK
	CREATED      = http.StatusCreated
	SERVER_ERROR = http.StatusInternalServerError
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as a the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 internal server error
	// response to the user. Note that we use the net/http constant
	// http.StatusInternalServerError here instead of the int 500 directly. Notice
	// that we use ... to pass the contents // of the files slice as variadic arguments.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler is now a method against the application struct
		// it can access its fields, including the structured logger. We'll
		// use this to create a log entry at Error level containing the error
		// message, also including the request method and URI as attributes to
		// assist with debugging
		app.serverError(w, r, err)
		http.Error(w, "Internal Server Error", SERVER_ERROR)
		return
	}

	// Then we use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal Server Error", SERVER_ERROR)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	if r.Method != http.MethodPost {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.

		// Use the Header().Set() method to add an 'Allow: POST' header to the // response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", NOT_ALLOWED)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

// Add a snippetCreatePost handler funciton
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Use the w.WriteHeader() method to send a 201 status code
	w.WriteHeader(CREATED)

	w.Write([]byte("Save a new snippet..."))

}
