package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route PUT /products/{id} updateProduct
// Updates a product
// responses:
//	201: noContent

// Update updates a product in the data store
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] getting product from context")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("[DEBUG] updating record with id", prod.ID)

	err := data.UpdateProduct(prod)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
