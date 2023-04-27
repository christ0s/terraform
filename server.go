package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    content, _ := ioutil.ReadFile("index.html")
    fmt.Fprintf(w, string(content))
  })

  fmt.Println("Server is up and running! Listening on port 8080.")
  http.ListenAndServe(":8080", nil)
}
