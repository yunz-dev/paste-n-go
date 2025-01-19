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

type application struct {
  errLog *log.Logger
  infoLog *log.Logger
}

func main() {

  cfg := new(Config)
  flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


  app := &application {
    errLog: errLog,
    infoLog: infoLog,
  }

  srv := &http.Server{
    Addr: cfg.Addr,
    ErrorLog: errLog,
    Handler: app.routes(),
  }


  // NOTE: ListenAndServe follows host:port
  // if no host then it will listen all all
  // hosts available on the computer
  infoLog.Printf("Starting Server on port %s", cfg.Addr)
  // initialise web server
  err := srv.ListenAndServe()
  errLog.Fatal(err)

}
