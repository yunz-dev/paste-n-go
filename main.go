package main

import (
  "log"
  "net/http"
)

// handler for home '/'
func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello from Paste n Go"))
}

func main() {
  // initialise new router ServeMux
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)

  log.Println("Starting Server on port :4000")
  // initialise web server
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)

}
