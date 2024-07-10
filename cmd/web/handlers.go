package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Galbeyte1/snippetbox/internal/models"
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

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, OK, "home.tmpl", data)
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, OK, "view.tmpl", data)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// THE FOLLOWING IS UNNECESSARY IN GO v1.22
	// // Use r.Method to check whether the request is using POST or not.
	// if r.Method != http.MethodPost {
	// 	// If it's not, use the w.WriteHeader() method to send a 405 status
	// 	// code and the w.Write() method to write a "Method Not Allowed"
	// 	// response body. We then return from the function so that the
	// 	// subsequent code is not executed.

	// 	// Use the Header().Set() method to add an 'Allow: POST' header to the // response header map. The first parameter is the header name, and
	// 	// the second parameter is the header value.
	// 	w.Header().Set("Allow", "POST")
	// 	http.Error(w, "Method Not Allowed", NOT_ALLOWED)
	// 	return
	// }
	w.Write([]byte("Create a new snippet..."))
}

// Add a snippetCreatePost handler funciton
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimbe Mount Fuji, \nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// // Use the w.WriteHeader() method to send a 201 status code
	// w.WriteHeader(CREATED)

	// w.Write([]byte("Save a new snippet..."))
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
