package main

import (
  "log"
  "net/http"
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
  w.Write([]byte("Showing Snippet..."))
}

// TODO: implement
func addSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    w.WriteHeader(405)
    w.Write([]byte("405 method not allowed"))
    return
  }
  w.Write([]byte("Adding Snippet..."))
}

func main() {
  // initialise new router ServeMux
  mux := http.NewServeMux()
  // add handler functions
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/add", addSnippet)

  // NOTE: ListenAndServe follows host:port
  // if no host then it will listen all all
  // hosts available on the computer
  log.Println("Starting Server on port :4000")
  // initialise web server
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)

}
