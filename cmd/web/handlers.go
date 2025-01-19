package main

import (
  "fmt"
  "net/http"
  "strconv"
  "html/template"
)

// handler for home '/'
func (app *application) home(w http.ResponseWriter, r *http.Request) {

  if r.URL.Path != "/" {
    app.notFound(w)
    return
    }

  files := []string{
    "./ui/html/home.page.tmpl",
    "./ui/html/base.layout.tmpl",
    "./ui/html/footer.partial.tmpl",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil {
    app.serverError(w, err)
    return
  }

  err = ts.Execute(w, nil)
  if err != nil {
    app.serverError(w, err)
  }
}

// TODO: implement
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(w)
    return
  }
  fmt.Fprintf(w, "Displaying snippet for id: %d: ", id)
}

// TODO: implement
func (app *application) addSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    app.clientError(w, http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Adding Snippet..."))
}
