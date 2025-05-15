package product

import (
	"fmt"
)

// Product contained in a store, such as flowers, vases, toys
type Product struct {
	id       int
	name     string
	price    float32
	quantity int
	discount float32
}

func NewProduct(id int, name string, price float32, quantity int, discount float32) *Product {
	return &Product{
		id:       id,
		name:     name,
		price:    price,
		quantity: quantity,
		discount: discount,
	}
}

func (p Product) GetName() string {
	return p.name
}

func (p *Product) SetPrice(price float32) {
	p.price = price
}

func (p Product) GetPrice() float32 {
	return p.price
}

func (p *Product) SetQuantity(quantity int) {
	p.quantity = quantity
}

func (p Product) GetQuantity() int {
	return p.quantity
}

func (p *Product) SaveItemLog() {
	fmt.Println("The Product \"" + p.name + "\"  was saved")
}
