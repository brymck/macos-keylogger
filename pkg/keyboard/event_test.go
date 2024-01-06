package keyboard

import (
	"fmt"
	"testing"
)

func TestNewEvent(t *testing.T) {
	var tests = []struct {
		key              string
		state            State
		ctrl             bool
		opt              bool
		shift            bool
		cmd              bool
		expectedNotation string
	}{
		{"a", Down, false, false, false, false, "a"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.expectedNotation)
		t.Run(testname, func(t *testing.T) {
			event := NewEvent(tt.key, tt.state, tt.ctrl, tt.opt, tt.shift, tt.cmd)
			actual := event.Notation
			expected := tt.expectedNotation
			if actual != expected {
				t.Errorf("got %s, want %s", actual, expected)
			}
		})
	}
}
