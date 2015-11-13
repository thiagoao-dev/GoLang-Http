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
var UserGetParameters = []Param{ { Name: "Email", ParametersDetails: Detail{ Type: "string", Description: "Get user Email", Required: true } } }

type DocMethod interface {}

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
    //w.Header().Set("Allow","DELETE,GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Allow","OPTIONS,GET")
	UserDoc := []DocMethod{}
	UserDoc = append(UserDoc, UserOPTIONS)
}

func Users(w http.ResponseWriter, r *http.Request) {
  
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