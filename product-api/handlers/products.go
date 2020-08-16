package handlers

import (
	"log"
	"net/http"
	"encoding/json"

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
	products := data.GetProducts()
	json, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
	}
	rw.Write(json)
}
