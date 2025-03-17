package cart

//Contains cart statuses with definitions
type CartStatus struct {
	id          int8
	name        string
	description string
}

func NewCartStatus(name string, description string) CartStatus {
	return CartStatus{
		name:        name,
		description: description,
	}
}

func (cs *CartStatus) GetId() int8 {
	return cs.id
}

func (cs *CartStatus) SetName(statusName string) {
	cs.name = statusName
}

func (cs *CartStatus) GetName() string {
	return cs.name
}

func (cs *CartStatus) SetDescription(statusDescription string) {
	cs.description = statusDescription
}

func (cs *CartStatus) GetDescription() string {
	return cs.description
}
