package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
  "encoding/json"
)

type Person struct {
  Name  string    `json="name"`
  Birth time.Time `json="birth"`
  Email string    `json="email"`
}

func (p *Person) Age() int {
  return 1
}

func main() {

  http.HandleFunc("/", Hello)
  http.HandleFunc("/user", GetUser)

  s := &http.Server{
    Addr:           ":8080",
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  log.Fatal(s.ListenAndServe())
}

// Hello
func Hello (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
  
  user, err := json.Marshal(Person{ Name:"Thiago Augustus de Oliveira", Birth: time.Date(1983, time.June, 9, 23, 0, 0, 0, time.UTC), Email:"thiagoaugustusdeoliveira@gmail.com"})
  
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "%s", user)
}