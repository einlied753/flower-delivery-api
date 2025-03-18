package cart

//Contains cart statuses
type CartStatus int

const (
	Empty CartStatus = iota
	InProgress
	Done
)
