package main

import (
  "fmt"
  "net/http"
  "strconv"
)

// handler for home '/'
func home(w http.ResponseWriter, r *http.Request) {

  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
    }

  w.Write([]byte("Hello from Paste n Go"))
}

// TODO: implement
func showSnippet(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, "Displaying snippet for id: %d: ", id)
}

// TODO: implement
func addSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    http.Error(w, "Method Not Allowed", 405)
    // w.WriteHeader(405)
    // w.Write([]byte("405 method not allowed"))
    return
  }
  w.Write([]byte("Adding Snippet..."))
}
