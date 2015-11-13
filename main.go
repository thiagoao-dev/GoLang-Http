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

type MethodGET struct {
  GET EndPoint
}

type MethodOPTIONS struct {
  OPTIONS EndPoint
}

type EndPoint struct {
  Description string  `json:"description"`
  Parameters  []Param `json:"parameters"`
}

type Param struct {
  Name              string `json:"name"`
  ParametersDetails Detail `json:"details"`
}

type Detail struct {
  Type        string `json:"type"`
  Description string `json:"description"`
  Required    bool   `json:"required"`
}

var UserOPTIONS = MethodOPTIONS{ OPTIONS: EndPoint{ Description: "User page" } }
// var UserGetParameters = []Param{ { Name: "Email" } }

func main() {

  http.HandleFunc("/api", Hello)
  http.HandleFunc("/api/users", Users)

  s := &http.Server{
    Addr:           ":8080",
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  log.Fatal(s.ListenAndServe())
}

func Hello(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
    //w.Header().Set("Allow","DELETE,GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Allow","OPTIONS,GET")
}

func Users(w http.ResponseWriter, r *http.Request) {
=======
>>>>>>> 8cb179b3fd32b845cc7a8c2a0a128c015c98e7cf
  
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