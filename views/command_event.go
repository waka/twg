package views

type CommandEvent struct {
	eventType CommandEventType
	value     []byte
}

func NewCommandEvent(eventType CommandEventType, value []byte) *CommandEvent {
	return &CommandEvent{eventType: eventType, value: value}
}

type CommandEventType int

const (
	CommandStart = iota + 1
	CommandEnd
	CommandAdd
	CommandDelete
	CommandLeft
	CommandRight
	CommandExecute
)

func (eventType CommandEventType) String() string {
	switch eventType {
	case CommandStart:
		return "CommandStart"
	case CommandEnd:
		return "CommandEnd"
	case CommandAdd:
		return "CommandAdd"
	case CommandDelete:
		return "CommandDelete"
	case CommandLeft:
		return "CommandLeft"
	case CommandRight:
		return "CommandRight"
	case CommandExecute:
		return "CommandExecute"
	}
	return "Undef"
}
