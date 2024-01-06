package keylogger

// #cgo LDFLAGS: -framework CoreFoundation -framework CoreGraphics
// #include "keylogger.c"
import "C"

import (
	"errors"
	"fmt"
	"os/user"

	"github.com/brymck/macos-keylogger/pkg/keyboard"
)

type Callback func(event keyboard.Event)

type cCallback func(keyCode C.int, stateCode C.State, ctrl C.bool, opt C.bool, shift C.bool, cmd C.bool)

type KeyLogger struct{}

var _f cCallback

//export handleButtonEvent
func handleButtonEvent(keyCode C.int, stateCode C.State, ctrl C.bool, opt C.bool, shift C.bool, cmd C.bool) {
	_f(keyCode, stateCode, ctrl, opt, shift, cmd)
}

func New() (*KeyLogger, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	if u.Uid != "0" {
		return nil, errors.New("must be run as root")
	}

	return &KeyLogger{}, nil
}

func (k *KeyLogger) Listen(f Callback) {
	_f = func(keyCode C.int, stateCode C.State, ctrl C.bool, opt C.bool, shift C.bool, cmd C.bool) {
		key, err := keyboard.ConvertKeyCode(int(keyCode))
		if err != nil {
			fmt.Printf("Could not convert key code: %d\n", int(keyCode))
			return
		}

		state := keyboard.State(stateCode)

		event := keyboard.NewEvent(key, state, bool(ctrl), bool(opt), bool(shift), bool(cmd))

		f(event)
	}
	C.listen()
}
