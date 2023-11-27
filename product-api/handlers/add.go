package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route POST /products products createProduct
// Adds a product to the database
// responses:
// 	204: created

// AddProduct adds a product to the data store
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.logger.Printf("Prod: %#v", product)
	data.AddProduct(product)
}
