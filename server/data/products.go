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
	"pirsch-phone-1": {
		Slug:        "pirsch-phone-1",
		Name:        "Pirsch Phone 1",
		Price:       "$799",
		Description: "Introducing the Pirsch Phone 1, a smartphone designed for those who value privacy and uncompromising build quality. Encased in a sleek aluminum body, this device combines modern aesthetics with robust durability, ensuring it stands the test of time.",
		Highlights:  "The Pirsch Phone 1 stands out with its steadfast commitment to user privacy and exceptional build quality. Crafted from high-grade aluminum, it’s as sturdy as it is sleek, offering a premium feel that lasts.",
		Details:     "With advanced privacy features like hardware-based encryption and a custom, privacy-focused OS, your data remains secure and under your control.",
		Img1:        "pirsch-phone-1/Pirsch Phone 6.png",
		Img2:        "pirsch-phone-1/Pirsch Phone 3.png",
	},
	"pirsch-pad-1": {
		Slug:        "pirsch-pad-1",
		Name:        "Pirsch Pad 1",
		Price:       "$1,299",
		Description: "Introducing the Pirsch Pad 1, a tablet that redefines the balance between privacy, performance, and premium build quality. Encased in a sleek aluminum shell, this tablet is designed to be both durable and aesthetically pleasing.",
		Highlights:  "The Pirsch Pad 1 combines the elegance of a finely crafted aluminum body with the assurance of cutting-edge privacy features. It’s built to last, with a design that’s both sturdy and sophisticated. With hardware-based encryption and a custom, privacy-oriented OS, your data is securely protected, giving you peace of mind.",
		Details:     "This tablet is equipped with advanced privacy features, including hardware-based encryption, a secure boot process, and a custom operating system focused on protecting your data. The Pirsch Pad 1 also boasts a high-resolution display, powerful processing capabilities, and extended battery life, ensuring it meets all your needs while keeping your information secure.",
		Img1:        "pirsch-pad-1/Pirsch Pad 6.png",
		Img2:        "pirsch-pad-1/Pirsch Pad 3.png",
	},
	"pirsch-watch-1": {
		Slug:        "pirsch-watch-1",
		Name:        "Pirsch Watch 1",
		Price:       "$399",
		Description: "Introducing the Pirsch Watch 1, a smartwatch that seamlessly blends privacy, durability, and sophisticated design. Housed in a sleek aluminum case, this smartwatch is built to withstand the rigors of daily life while maintaining a refined, modern appearance.",
		Highlights:  "Its combination of robust build quality and dedicated privacy features. Crafted from high-grade aluminum, it offers both strength and elegance in a compact form. With hardware-based encryption and a privacy-focused operating system, your personal data stays secure, giving you the confidence to use your smartwatch freely.",
		Details:     "The Pirsch Watch 1 is meticulously engineered with a durable aluminum casing, ensuring that it is both lightweight and resilient. This smartwatch includes a comprehensive suite of privacy features, such as hardware-based encryption, secure boot, and a custom OS that is designed to protect your data at all times.",
		Img1:        "pirsch-watch-1/Pirsch Watch 1.png",
		Img2:        "pirsch-watch-1/Pirsch Watch 4.png",
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
