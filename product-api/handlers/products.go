package handlers

import (
	"log"
	"fmt"
	"encoding/json"
	"net/http"

	"product-api/data"
)

// Products the products handlers
type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (handler *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProducts(rw, r)
	case "PUT":
		putProducts(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func putProducts(rw http.ResponseWriter, r *http.Request) {
	var product data.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
	data.PutProducts(product)
}
