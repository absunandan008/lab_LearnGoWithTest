package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

const finalWord = "Go!"
const countdownStart = 3

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
