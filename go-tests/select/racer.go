package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	durationA := getDuration(a)
	durationB := getDuration(b)

	if durationA > durationB {
		return b
	} else {
		return a
	}
}

func getDuration(url string) time.Duration {
	timer := time.Now()
	//fmt.Printf("Calling %q", url)
	http.Get(url)
	return time.Since(timer)
}

/*
The select statement is used for non-blocking communication on multiple channels. It waits for the first of several channel operations to become ready.
The select statement is non-deterministic if both channels become ready at the same time — meaning either a or b could be returned in that case.
Even though ping(a) is called before ping(b), the actual network requests (http.Get(...)) happen concurrently.
Since both goroutines are doing their HTTP GETs concurrently, the select will return the channel that closes first — that is, the URL that responds faster
*/
func RacerOptimised(a, b string) (winner string) {

	select {
	case <-ping2(a):
		return a
	case <-ping2(b):
		return b
	}
}

/*
Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation
versus a bool. Since we are closing and not sending anything on the chan, why allocate anything
*/

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func ping2(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
