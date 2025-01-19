package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
  Addr string
}

func main() {

  cfg := new(Config)
  flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
  flag.Parse()
  fileServer := http.FileServer(http.Dir("./ui/static"))

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  // initialise new router ServeMux
  mux := http.NewServeMux()

  srv := &http.Server{
    Addr: cfg.Addr,
    ErrorLog: errLog,
    Handler: mux,
  }
  // add handler functions
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/add", addSnippet)
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  // NOTE: ListenAndServe follows host:port
  // if no host then it will listen all all
  // hosts available on the computer
  infoLog.Printf("Starting Server on port %s", cfg.Addr)
  // initialise web server
  err := srv.ListenAndServe()
  errLog.Fatal(err)

}
