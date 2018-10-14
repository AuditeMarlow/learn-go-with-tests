package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, err := fmt.Fprintln(writer, i)

		panicOnError(err)

		sleeper.Sleep()
	}

	_, err := fmt.Fprint(writer, finalWord)

	panicOnError(err)
}

func panicOnError(err error) {
	if err != nil {
		panic("could not write to output")
	}
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
