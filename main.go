package main

import (
	"fmt"

	"github.com/brymck/macos-keylogger/pkg/keyboard"
	"github.com/brymck/macos-keylogger/pkg/keylogger"
)

func main() {
	kl, err := keylogger.New()
	if err != nil {
		panic(err)
	}
	kl.Listen(func(event keyboard.Event) {
		if event.State == keyboard.Down {
			fmt.Printf("%s %v\n", event.Notation, event)
		}
	})
}
