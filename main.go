package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
)

func main() {

  http.HandleFunc("/", Hello)


  s := &http.Server{
    Addr:           ":8080",
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  log.Fatal(s.ListenAndServe())
}


func Hello (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}