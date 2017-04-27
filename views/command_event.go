package views

type CommandEvent struct {
	eventType CommandEventType
}

func NewCommandEvent(eventType CommandEventType) *CommandEvent {
	return &CommandEvent{eventType: eventType}
}

type CommandEventType int

const (
	CommandAdd = iota + 1
	CommandDelete
	CommandLeft
	CommandRight
)
