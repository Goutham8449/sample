package data

import "time"

type Product struct {
	ID        int
	Name      string
	Desc      string
	Price     float32
	SKU       string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Latte",
		Desc:      "Frothy milky coffee",
		Price:     2.45,
		SKU:       "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Short and strong coffee",
		Price:     1.99,
		SKU:       "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
