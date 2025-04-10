package order

import (
	"fmt"
	"strconv"
	"time"
)

// Link between orders and products, contains products added to order
type OrderProduct struct {
	id           int
	orderId      int
	productId    int
	productCount int
}

func NewOrderProduct(orderId int, productId int, productCount int) *OrderProduct {
	return &OrderProduct{
		id:           1,
		orderId:      orderId,
		productId:    productId,
		productCount: productCount,
	}
}

func (op *OrderProduct) GetOrderId() int {
	return op.orderId
}

func (op *OrderProduct) GetProductId() int {
	return op.productId
}

func (op *OrderProduct) SetProductCount(productCount int) {
	op.productCount = productCount
}

func (op *OrderProduct) GetProductCount() int {
	return op.productCount
}

func (op *OrderProduct) SaveItem() {
	time.Sleep(time.Second)
	fmt.Println("The OrderProduct " + strconv.Itoa(op.id) + " was saved")
}
