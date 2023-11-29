package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// responses:
// 	201: noContent
// 	404: errorResponse
//  501: errorResponse

// DeleteProduct handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Println("[DEBUG] getting id from url")

	id := getProductID(r)
	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record, id does not exist")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		// http.Error(w, "Product not found", http.StatusInternalServerError)
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
