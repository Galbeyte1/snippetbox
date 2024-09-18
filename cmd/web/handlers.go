package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Galbeyte1/snippetbox/internal/models"
	"github.com/Galbeyte1/snippetbox/internal/validator"
)

const (
	NOT_ALLOWED   = http.StatusMethodNotAllowed
	OK            = http.StatusOK
	CREATED       = http.StatusCreated
	SERVER_ERROR  = http.StatusInternalServerError
	SEE_OTHER     = http.StatusSeeOther
	BAD_REQUEST   = http.StatusBadRequest
	UNPROCESSABLE = http.StatusUnprocessableEntity
)

type snippetCreateForm struct {
	Title               string
	Content             string
	Expires             int
	validator.Validator `form:"-"`
}

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

	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, OK, "create.tmpl", data)
}

// Add a snippetCreatePost handler funciton
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	var form snippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, BAD_REQUEST)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Title), "content", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 1000), "content", "This field cannot be more than 1000 characters long")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "This field must be equal 1, 7, or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, UNPROCESSABLE, "create.tmpl", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), SEE_OTHER)
}
