package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type DefaultSleeper struct {
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		//there is a dependency on sleep for this func. So even test methods will have to sleep. This results in longer test runs
		// we can avoid this by using DI and mocking, giving control from outside. In tests, we can control how long to sleep
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
