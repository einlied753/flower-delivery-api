package cart

import (
	"api/internal/model"
	"fmt"
	"strconv"
	"time"
)

// Link between cart and products, showing what products are in the cart
type CartProduct struct {
	id           int
	cartId       int
	productId    int
	productCount int
}

func NewCartProduct(id int, cartId int, productId int, productCount int) model.Item {
	return CartProduct{
		id:           id,
		cartId:       cartId,
		productId:    productId,
		productCount: productCount,
	}
}

func (cp *CartProduct) SetProductCount(productCount int) {
	cp.productCount = productCount
}

func (cp *CartProduct) GetProductCount() int {
	return cp.productCount
}

func (cp CartProduct) SaveItem() {
	time.Sleep(time.Second)
	fmt.Println("The CartProduct " + strconv.Itoa(cp.id) + " was saved")
}
