package order

import (
	"fmt"
	"strconv"
	"time"
)

// Order created by user after selecting products
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

func NewOrder(userId int, cost float32, productCount int, email string, address string, phone string) *Order {
	return &Order{
		id:           1,
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

func (o Order) GetId() int {
	return o.id
}

func (o Order) GetUserId() int {
	return o.userId
}

func (o *Order) SetStatus(orderStatus OrderStatus) {
	o.status = orderStatus
}

func (o Order) GetStatus() OrderStatus {
	return o.status
}

func (o Order) GetCost() float32 {
	return o.cost
}

func (o Order) GetProductCount() int {
	return o.productCount
}

func (o Order) GetCreated() string {
	return o.created
}

func (o *Order) SetEmail(email string) {
	o.email = email
}

func (o Order) GetEmail() string {
	return o.email
}

func (o *Order) SetAddress(address string) {
	o.address = address
}

func (o Order) GetAddress() string {
	return o.address
}

func (o Order) GetPhone() string {
	return o.phone
}

func (o *Order) SaveItemLog() {
	fmt.Println("The Order " + strconv.Itoa(o.id) + " was saved")
}

func (o Order) GetStringOrderStatus() string {
	return o.status.GetStringName()
}
