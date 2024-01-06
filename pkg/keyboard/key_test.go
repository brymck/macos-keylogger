package keyboard

import (
	"fmt"
	"testing"
)

func TestConvertKeyCode(t *testing.T) {
	var tests = []struct {
		keyCode int
		wantKey rune
		wantErr error
	}{
		{-1, ' ', ErrInvalidKeyCode},
		{0, 'a', nil},
		{10, ' ', ErrInvalidKeyCode},
		{126, '↑', nil},
		{127, ' ', ErrInvalidKeyCode},
		{128, ' ', ErrInvalidKeyCode},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.keyCode)
		t.Run(testname, func(t *testing.T) {
			key, err := ConvertKeyCode(tt.keyCode)
			if key != tt.wantKey {
				t.Errorf("got %c, want %c", key, tt.wantKey)
			}
			if err != tt.wantErr {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}
