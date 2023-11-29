package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route POST /products products createProduct
// Adds a product to the database
// responses:
// 	204: created
//	422: errorValidation
//  501: errorResponse

// Create adds a product to the data store
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product : %#v\n", product)
	data.AddProduct(product)
}
