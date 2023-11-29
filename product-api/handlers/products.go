package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

type KeyProduct struct{}

type GenericError struct {
	Message string `json:"message"`
}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		// http.Error(w, "Unable to convert id", http.StatusBadRequest)
		panic(err)
	}

	return id
}

type ValidationError struct {
	Messages []string `json:"messages"`
}
