package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello First Handler
type Hello struct {
	logger *log.Logger
}

// NewHello creates an instance of the Hello handler
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (handler *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handler.logger.Println("Hello World From the Hello Handler!")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello %s\n", data)
}
