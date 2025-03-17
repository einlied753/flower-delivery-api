package order

//Contains order statuses with definitions
type OrderStatus struct {
	id          int8
	name        string
	description string
}

func NewOrderStatus(name string, description string) OrderStatus {
	return OrderStatus{
		name:        name,
		description: description,
	}
}

func (os *OrderStatus) GetId() int8 {
	return os.id
}

func (os *OrderStatus) SetName(statusName string) {
	os.name = statusName
}

func (os *OrderStatus) GetName() string {
	return os.name
}

func (os *OrderStatus) SetDescription(statusDescription string) {
	os.description = statusDescription
}

func (os *OrderStatus) GetDescription() string {
	return os.description
}
