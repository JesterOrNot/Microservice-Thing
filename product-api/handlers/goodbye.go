package handlers

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

// Goodbye The Goodbye handler
type Goodbye struct {
	logger *log.Logger
}

// NewGoodbye creates an instance of the goodbye handler
func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (handler *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handler.logger.Println("Goodbye World From The GoodBye Handler")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Goodbye %s\n", data)
}
