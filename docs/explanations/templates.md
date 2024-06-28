# Templates

As I add more pages to the web application, there will be some shared, boilerplate, HTML markup that I want to include on every page — like the header, navigation and metadata inside the <head> HTML element.

To prevent duplication and save typing, it’s a good idea to create a `base` (or master) template which contains this shared content, which I can then compose with the page-specific markup for the individual pages.

It's important to note that the file containing the base template must be the **_first_** file in the slice.

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
that we use ... to pass the contents of the files slice as variadic arguments.

```Go
ts, err := template.ParseFiles(files...)
```

Use the ExecuteTemplate() method to write the content of the "base"
template as the response body.

Instead of containing HTML directly, our template set contains 3 named templates
— `base`, `title` and `main`. We use the `ExecuteTemplate()` method to tell Go that we specifically want to respond using the content of the base template (which in turn invokes our title and main templates).

```Go
err = ts.ExecuteTemplate(w, "base", nil)
if err != nil {
	...
}
```
