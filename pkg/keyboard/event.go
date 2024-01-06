package keyboard

import (
	"strings"
	"unicode"
)

type Event struct {
	Notation string
	Ch       rune
	Key      string
	State    State
	Ctrl     bool
	Opt      bool
	Shift    bool
	Cmd      bool
}

func NewEvent(key string, ch rune, state State, ctrl bool, opt bool, shift bool, cmd bool) Event {
	var builder strings.Builder

	if !unicode.IsPrint(ch) || ch == ' ' {
		ch = 0
	}

	if ctrl {
		builder.WriteString("C-")
	}

	if opt {
		builder.WriteString("A-")
	}

	if shift {
		if ch == 0 {
			builder.WriteString("S-")
		}
	}

	if cmd {
		builder.WriteString("D-")
	}

	if ch == 0 {
		builder.WriteString(key)
	} else {
		builder.WriteRune(ch)
	}

	notation := builder.String()
	if len(notation) > 1 {
		notation = "<" + notation + ">"
	}

	return Event{
		Notation: notation,
		Ch:       ch,
		Key:      key,
		State:    state,
		Ctrl:     ctrl,
		Opt:      opt,
		Shift:    shift,
		Cmd:      cmd,
	}
}
