package main

import (
	"log"
	"fmt"
	"io/ioutil"
  "net/http"
)

func main() {
  http.HandleFunc("/greet", func(rw http.ResponseWriter, req *http.Request) {
    data, err := ioutil.ReadAll(req.Body)
    if err != nil {
      fmt.Println("Error: \033[1;31mRead error\033[m")
      rw.WriteHeader(http.StatusBadRequest)
    }
    log.Println(req.Method, req.RequestURI)
    fmt.Fprintf(rw, "Hello %s", data)
  })
  fmt.Println("Server listening on Port 8000")
  http.ListenAndServe(":8000", nil)
}
