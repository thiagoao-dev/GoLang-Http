package main

import (
  "fmt"
  "log"
  "time"
  //"strconv"
  "net/http"
  "encoding/json"
)

type Person struct {
  Name  string    `json="name"`
  Birth time.Time `json="birth"`
  Email string    `json="email"`
  Age   int       `json="age"`
}

func (p *Person) SetAge() {
  p.Age = time.Now().Year() - p.Birth.Year()
}

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

func Hello(w http.ResponseWriter, r *http.Request) {
  
  user := Person{ Name:"Jon Jones", Birth: time.Date(1987, time.July, 19, 0, 0, 0, 0, time.UTC), Email:"j.jones@ufc.com"}
  user.SetAge()
  
  userJson, err := json.Marshal(user)
  
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "%s", userJson)
}