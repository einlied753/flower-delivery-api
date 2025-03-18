package order

import "time"

//Order created by user after selecting products
type Order struct {
	id           int
	userId       int
	status       OrderStatus
	cost         float32
	productCount int
	created      string
	email        string
	address      string
	phone        string
}

func NewOrder(userId int, cost float32, productCount int, email string, address string, phone string) Order {
	return Order{
		userId:       userId,
		status:       New,
		cost:         cost,
		productCount: productCount,
		created:      time.Now().GoString(),
		email:        email,
		address:      address,
		phone:        phone,
	}
}

func (o *Order) SetStatus(orderStatus OrderStatus) {
	o.status = orderStatus
}

func (o *Order) GetStatus() OrderStatus {
	return o.status
}

func (o *Order) SetEmail(email string) {
	o.email = email
}

func (o *Order) GetEmail() string {
	return o.email
}

func (o *Order) SetAddress(address string) {
	o.address = address
}

func (o *Order) GetAddress() string {
	return o.address
}
