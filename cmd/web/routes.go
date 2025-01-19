package main

import "net/http"

func (app *application) routes() *http.ServeMux {

  // initialise new router ServeMux
  mux := http.NewServeMux()

  // add handler functions
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippet", app.showSnippet)
  mux.HandleFunc("/snippet/add", app.addSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  return mux
}
