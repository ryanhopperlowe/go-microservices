package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.logger.Printf("Prod: %#v", product)
	data.AddProduct(product)
}

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

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading Product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.logger.Println("[ERROR] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating Product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
