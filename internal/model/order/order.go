package order

import "time"

//Order created by user after selecting products
type Order struct {
	id           int
	userId       int
	statusId     int8
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
		statusId:     1, //draft
		cost:         cost,
		productCount: productCount,
		created:      time.Now().GoString(),
		email:        email,
		address:      address,
		phone:        phone,
	}
}

func (o *Order) SetStatus(orderStatusId int8) {
	o.statusId = orderStatusId
}

func (o *Order) GetStatus() int8 {
	return o.statusId
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
