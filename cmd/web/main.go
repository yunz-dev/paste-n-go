package main

import(
  "flag"
  "log"
  "net/http"
)

type Config struct {
  Addr string
}

func main() {

  cfg := new(Config)
  flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
  flag.Parse()
  fileServer := http.FileServer(http.Dir("./ui/static"))
  // initialise new router ServeMux
  mux := http.NewServeMux()
  // add handler functions
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/add", addSnippet)
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  // NOTE: ListenAndServe follows host:port
  // if no host then it will listen all all
  // hosts available on the computer
  log.Printf("Starting Server on port %s", cfg.Addr)
  // initialise web server
  err := http.ListenAndServe(cfg.Addr, mux)
  log.Fatal(err)

}
