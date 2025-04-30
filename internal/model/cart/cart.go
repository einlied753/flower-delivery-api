package cart

import (
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

func NewCart(userId int, orderId int) *Cart {
	return &Cart{
		id:      1,
		userId:  userId,
		orderId: orderId,
		status:  Empty,
		created: time.Now().GoString(),
	}
}

func (c Cart) GetId() int {
	return c.id
}

func (c *Cart) SetStatus(cartStatus CartStatus) {
	c.status = cartStatus
}

func (c Cart) GetStatus() CartStatus {
	return c.status
}

func (c *Cart) SaveItem() {
	time.Sleep(time.Second)
	//fmt.Println("The Cart " + strconv.Itoa(c.id) + " was saved")
}
