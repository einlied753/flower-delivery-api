package cart

import (
	"fmt"
	"strconv"
)

// Link between cart and products, showing what products are in the cart
type CartProduct struct {
	id           int
	cartId       int
	productId    int
	productCount int
}

func NewCartProduct(id int, cartId int, productId int, productCount int) *CartProduct {
	return &CartProduct{
		id:           id,
		cartId:       cartId,
		productId:    productId,
		productCount: productCount,
	}
}

func (cp CartProduct) GetId() int {
	return cp.id
}

func (cp CartProduct) GetCartId() int {
	return cp.cartId
}

func (cp CartProduct) GetProductId() int {
	return cp.productId
}

func (cp *CartProduct) SetProductCount(productCount int) {
	cp.productCount = productCount
}

func (cp CartProduct) GetProductCount() int {
	return cp.productCount
}

func (cp *CartProduct) SaveItemLog() {
	fmt.Println("The CartProduct " + strconv.Itoa(cp.id) + " was saved")
}
