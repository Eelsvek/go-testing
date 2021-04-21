package main

import (
	"fmt"
	"io"
	"os"
)

const countdownStart = 3
const finalWord = "Go!"

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprint(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
