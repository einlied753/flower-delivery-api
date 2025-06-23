package cart

import (
	"fmt"
	"strconv"
	"time"
)

// Cart for creating order
type Cart struct {
	id      int
	userId  int
	orderId int
	status  CartStatus
	created string
}

func NewCart(id int, userId int, orderId int) *Cart {
	return &Cart{
		id:      id,
		userId:  userId,
		orderId: orderId,
		status:  Empty,
		created: time.Now().GoString(),
	}
}

func (c Cart) GetId() int {
	return c.id
}

func (c Cart) GetUserId() int {
	return c.userId
}

func (c Cart) GetOrderId() int {
	return c.orderId
}

func (c *Cart) SetStatus(cartStatus CartStatus) {
	c.status = cartStatus
}

func (c Cart) GetStatus() CartStatus {
	return c.status
}

func (c Cart) GetCreated() string {
	return c.created
}

func (c *Cart) SaveItemLog() {
	fmt.Println("The Cart " + strconv.Itoa(c.id) + " was saved")
}

func (c Cart) GetStringCartStatus() string {
	return c.status.GetStringName()
}
