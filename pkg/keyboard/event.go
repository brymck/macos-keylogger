package keyboard

import "strings"

type Event struct {
	Notation string
	Key      string
	State    State
	Ctrl     bool
	Opt      bool
	Shift    bool
	Cmd      bool
}

func NewEvent(key string, state State, ctrl bool, opt bool, shift bool, cmd bool) Event {
	var builder strings.Builder

	if ctrl {
		builder.WriteString("C-")
	}

	if opt {
		builder.WriteString("A-")
	}

	if shift {
		builder.WriteString("S-")
	}

	if cmd {
		builder.WriteString("D-")
	}

	builder.WriteString(key)

	notation := builder.String()
	if len(notation) > 1 {
		notation = "<" + notation + ">"
	}

	return Event{
		Notation: notation,
		Key:      key,
		State:    state,
		Ctrl:     ctrl,
		Opt:      opt,
		Shift:    shift,
		Cmd:      cmd,
	}
}
