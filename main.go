package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/brymck/macos-keylogger/pkg/keyboard"
	"github.com/brymck/macos-keylogger/pkg/keylogger"
)

const defaultLingerTime = 500 * time.Millisecond

func main() {
	var lingerTime time.Duration

	if len(os.Args) > 1 {
		lingerMs, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error: linger time must be an integer\n")
		}
		lingerTime = time.Duration(lingerMs) * time.Millisecond
	} else {
		lingerTime = defaultLingerTime
	}

	kl, err := keylogger.New()
	if err != nil {
		panic(err)
	}

	last := time.Now()
	kl.Listen(func(event keyboard.Event) {
		if event.State == keyboard.Down {
			now := time.Now()
			if now.Sub(last) > lingerTime {
				fmt.Print("\n")
			}
			fmt.Print(event.Notation)
			last = now
		}
	})
}
