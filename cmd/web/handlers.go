package main

import (
  "fmt"
  "net/http"
  "strconv"
  "log"
  "html/template"
)

// handler for home '/'
func (app *application) home(w http.ResponseWriter, r *http.Request) {

  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
    }

  files := []string{
    "./ui/html/home.page.tmpl",
    "./ui/html/base.layout.tmpl",
    "./ui/html/footer.partial.tmpl",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil {
    app.errLog.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }

  err = ts.Execute(w, nil)
  if err != nil {
    app.errLog.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
  }
  // w.Write([]byte("Hello from Paste n Go"))
}

// TODO: implement
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, "Displaying snippet for id: %d: ", id)
}

// TODO: implement
func (app *application) addSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    http.Error(w, "Method Not Allowed", 405)
    return
  }
  w.Write([]byte("Adding Snippet..."))
}
