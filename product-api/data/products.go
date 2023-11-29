package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for this product
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// the name of this product
	//
	// required: true
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: true
	Description string `json:"description"`

	// the price for this product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for this product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`

	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func GetProducts() Products {
	return productList
}

func GetProductById(id int) (*Product, error) {
	index := findIndexByProductID(id)

	if index == -1 {
		return nil, ErrProductNotFound
	}

	return productList[index], nil
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	if i == len(productList)-1 {
		productList = productList[:i]
	} else {
		productList = append(productList[:i], productList[i+1])
	}

	return nil
}

func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func UpdateProduct(p *Product) error {
	_, i, err := findProduct(p.ID)

	if err != nil {
		return ErrProductNotFound
	}

	productList[i] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextId() int {
	lastPruduct := productList[len(productList)-1]
	return lastPruduct.ID + 1
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
