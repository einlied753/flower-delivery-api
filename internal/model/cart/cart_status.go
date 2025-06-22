package cart

//Contains cart statuses
type CartStatus int

const (
	Empty CartStatus = iota
	InProgress
	Done
)

func (cs CartStatus) GetStringName() string {
	switch cs {
	case Empty:
		return "Empty"
	case InProgress:
		return "InProgress"
	case Done:
		return "Done"
	default:
		return "Unknown status type"
	}
}
