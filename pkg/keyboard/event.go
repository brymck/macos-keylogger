package keyboard

import (
	"strings"
	"unicode"
)

type Event struct {
	Notation string
	Ch       rune
	Key      rune
	State    State
	Ctrl     bool
	Opt      bool
	Shift    bool
	Cmd      bool
}

func NewEvent(key rune, ch rune, state State, ctrl bool, opt bool, shift bool, cmd bool) Event {
	var builder strings.Builder

	if !unicode.IsPrint(ch) || ch == ' ' {
		ch = 0
	}

	if ctrl {
		builder.WriteRune('⌃')
	}

	if opt {
		builder.WriteRune('⌥')
	}

	if shift {
		if ch == 0 {
			builder.WriteRune('⇧')
		}
	}

	if cmd {
		builder.WriteRune('⌘')
	}

	if ch == 0 {
		builder.WriteRune(key)
	} else {
		builder.WriteRune(ch)
	}

	notation := builder.String()

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
