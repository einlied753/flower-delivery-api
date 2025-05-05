package order

//Contains order statuses
type OrderStatus int

const (
	New OrderStatus = iota
	Booked
	Paid
	Collecting
	Sent
	Delivered
)

var statusDescription = map[OrderStatus]string{
	New:        "new order was created",
	Booked:     "order booked",
	Paid:       "order paid",
	Collecting: "order is currently being collected",
	Sent:       "order sent by courier",
	Delivered:  "order delivered",
}

func (os OrderStatus) GetDescription() string {
	return statusDescription[os]
}
