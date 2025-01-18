package main

import(
  "log"
  "net/http"
)

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
