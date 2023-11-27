package handlers

import "product-api/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct{}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters createProduct updateProduct
type productParamsWrapper struct {
	// in: body
	// required: true
	Body data.Product
}

// swagger:response created
type productsCreated struct{}
