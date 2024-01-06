package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/brymck/macos-keylogger/pkg/keyboard"
	"github.com/brymck/macos-keylogger/pkg/keylogger"
)

const defaultLingerTime = 500 * time.Millisecond

func main() {
	var f io.Writer
	var lingerTime time.Duration

	if len(os.Args) < 1 {
		fmt.Printf("Usage: %s [corpus file]\n", os.Args[0])
		fmt.Printf("       %s [corpus file] [linger time in ms]\n", os.Args[0])
		os.Exit(1)
	} else if len(os.Args) > 2 {
		lingerMs, err := strconv.Atoi(os.Args[2])
		if err != nil || lingerMs < 0 {
			fmt.Printf("Error: linger time must be a non-negative integer\n")
			os.Exit(1)
		}
		lingerTime = time.Duration(lingerMs) * time.Millisecond
	} else {
		lingerTime = defaultLingerTime
	}

	// Either print to stdout if the first command line argument is "-" or
	// append to the specified  file
	corpusFile := os.Args[1]
	if corpusFile == "-" {
		f = os.Stdout
	} else {
		file, err := os.OpenFile(corpusFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		f = file
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
				fmt.Fprint(f, "\n")
			}

			fmt.Fprint(f, event.Notation)
			last = now
		}
	})
}
