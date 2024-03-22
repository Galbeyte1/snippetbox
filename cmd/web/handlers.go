package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	NOT_ALLOWED  = http.StatusMethodNotAllowed
	OK           = http.StatusOK
	SERVER_ERROR = http.StatusInternalServerError
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as a the response body.
func home(w http.ResponseWriter, r *http.Request) {
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
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", SERVER_ERROR)
		return
	}

	// Then we use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", SERVER_ERROR)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {

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
