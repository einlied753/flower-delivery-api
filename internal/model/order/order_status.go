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

func (os OrderStatus) GetStringName() string {
	switch os {
	case New:
		return "New"
	case Booked:
		return "Booked"
	case Paid:
		return "Paid"
	case Collecting:
		return "Collecting"
	case Sent:
		return "Sent"
	case Delivered:
		return "Delivered"
	default:
		return "Unknown status type"
	}
}
