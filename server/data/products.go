package data

import (
	"slices"
	"strings"
)

// Product is the data for a product.
type Product struct {
	Slug        string
	Name        string
	Price       string
	Description string
	Details     string
	Highlights  string
	Img1        string
	Img2        string
}

// Products is a map of all products with the slug as index.
var Products = map[string]Product{
	"device-a": {
		Slug:        "device-a",
		Name:        "Device A",
		Price:       "$195.98",
		Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Highlights:  "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.",
		Details:     "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea.",
		Img1:        "test.jpeg",
		Img2:        "test2.jpeg",
	},
}

// ListProducts returns a list of all products in alphabetical order.
func ListProducts() []Product {
	l := make([]Product, 0, len(Products))

	for _, v := range Products {
		l = append(l, v)
	}

	slices.SortFunc(l, func(a, b Product) int {
		return strings.Compare(a.Slug, b.Slug)
	})
	return l
}

// GetProduct returns a product for given slug or nil, if not found.
func GetProduct(slug string) *Product {
	p, ok := Products[slug]

	if !ok {
		return nil
	}

	return &p
}
