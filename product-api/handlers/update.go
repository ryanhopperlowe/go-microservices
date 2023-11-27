package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} updateProduct
// Updates a product
// responses:
//	201: noContent

// UpdateProducts updates a product in the data store
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.logger.Println("Handle PUT Product: ", id)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
