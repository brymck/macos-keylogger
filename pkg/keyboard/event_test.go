package keyboard

import "testing"

func TestNewEvent(t *testing.T) {
	var tests = []struct {
		key   rune
        ch    rune
		state State
		ctrl  bool
		opt   bool
		shift bool
		cmd   bool
		want  string
	}{
		{'a', 'a', Down, false, false, false, false, "a"},
		{'a', 'a', Down, true, false, false, false, "⌃a"},
		{'a', 'a', Down, false, true, false, false, "⌥a"},
		{'a', 'A', Down, false, false, true, false, "A"},
		{'a', 'a', Down, false, false, false, true, "⌘a"},
		{'a', 'A', Down, true, true, true, true, "⌃⌥⌘A"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			event := NewEvent(tt.key, tt.ch, tt.state, tt.ctrl, tt.opt, tt.shift, tt.cmd)
			actual := event.Notation
			expected := tt.want
			if actual != expected {
				t.Errorf("got %s, want %s", actual, expected)
			}
		})
	}
}
