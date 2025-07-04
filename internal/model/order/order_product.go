package order

import (
	"fmt"
	"strconv"
)

// Link between orders and products, contains products added to order
type OrderProduct struct {
	id           int
	orderId      int
	productId    int
	productCount int
}

func NewOrderProduct(id int, orderId int, productId int, productCount int) *OrderProduct {
	return &OrderProduct{
		id:           id,
		orderId:      orderId,
		productId:    productId,
		productCount: productCount,
	}
}

func (op OrderProduct) GetId() int {
	return op.id
}

func (op OrderProduct) GetOrderId() int {
	return op.orderId
}

func (op OrderProduct) GetProductId() int {
	return op.productId
}

func (op *OrderProduct) SetProductCount(productCount int) {
	op.productCount = productCount
}

func (op OrderProduct) GetProductCount() int {
	return op.productCount
}

func (op *OrderProduct) SaveItemLog() {
	fmt.Println("The OrderProduct " + strconv.Itoa(op.id) + " was saved")
}
