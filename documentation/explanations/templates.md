# Templating

It's important to note that the file containing our base template must be the **_first_** file in the slice.

```Go
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
```

Use the `template.ParseFiles()` function to read the template file into a
template set. If there's an error, we log the detailed error message and use
the `http.Error()` function to send a generic `500 internal server error`
response to the user. Note that we use the net/http constant
http.StatusInternalServerError here instead of the int 500 directly. Notice
that we use ... to pass the contents
of the files slice as variadic arguments.
