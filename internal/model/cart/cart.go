package cart

import "time"

//Cart for creating order
type Cart struct {
	id       int
	userId   int
	orderId  int
	statusId int8
	created  string
}

func NewCart(userId int, orderId int) Cart {
	return Cart{
		userId:   userId,
		orderId:  orderId,
		statusId: 1,
		created:  time.Now().GoString(),
	}
}

func (c *Cart) SetStatus(cartStatusId int8) {
	c.statusId = cartStatusId
}

func (c *Cart) GetStatus() int8 {
	return c.statusId
}
