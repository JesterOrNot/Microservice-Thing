package data

import (
	"time"
)

// Product The Product Serialization
type Product struct {
	ID   int
	Name string
	Description string
	Price float32
	SKU string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky Coffee",
		Price: 2.45,
		SKU: "latte1",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Frothy milky Coffee",
		Price: 2.45,
		SKU: "espresso1",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
