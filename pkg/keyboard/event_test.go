package keyboard

import "testing"

func TestNewEvent(t *testing.T) {
	var tests = []struct {
		key   string
		state State
		ctrl  bool
		opt   bool
		shift bool
		cmd   bool
		want  string
	}{
		{"a", Down, false, false, false, false, "a"},
		{"a", Down, true, false, false, false, "<C-a>"},
		{"a", Down, false, true, false, false, "<A-a>"},
		{"a", Down, false, false, true, false, "<S-a>"},
		{"a", Down, false, false, false, true, "<D-a>"},
		{"a", Down, true, true, true, true, "<C-A-S-D-a>"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			event := NewEvent(tt.key, tt.state, tt.ctrl, tt.opt, tt.shift, tt.cmd)
			actual := event.Notation
			expected := tt.want
			if actual != expected {
				t.Errorf("got %s, want %s", actual, expected)
			}
		})
	}
}
