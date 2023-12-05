// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
// swagger:meta
package handlers

import "product-api/data"

//
// Note: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response productResponse
type productResponseWrapper struct {
	// A product in the system
	// in: body
	Body data.Product
}

// swagger:response noContent
type productsNoContent struct{}

// swagger:parameters listSingleProduct deleteProduct
type productIDParameterWrapper struct {
	// The id of the product for which the operation relates
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
