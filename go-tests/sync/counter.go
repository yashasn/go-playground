package sync

import "sync"

/*
Go’s sync.Mutex is not strictly fair — it doesn't guarantee FIFO order of lock acquisition.
Instead, it is optimized for low contention and speed, favoring throughput.
When contention is high, it uses techniques like spinning (active waiting for a short time)
before sleeping the goroutine to improve performance.

Spinning, also known as busy waiting, is a technique where a thread (or goroutine)
repeatedly checks a condition in a tight loop without yielding the CPU.

TODO - https://medium.com/@cheshirysh/handmade-mutex-in-28-lines-of-go-205ff5b13dbc

When to use Mutex or Channel?
Mutexes are typically used for protecting shared data structures, while channels are used for communication between goroutines.
Mutexes are more efficient for low contention scenarios, while channels are better for high contention or complex synchronization patterns.
Mutexes are a low-level synchronization primitive that allows you to lock and unlock critical sections of code.
Channels are a higher-level abstraction that provides a way to communicate between goroutines and synchronize their execution.
Mutexes are generally faster than channels for simple locking scenarios, but channels can be more expressive and easier to reason about in complex concurrent programs.
*/

type Counter struct {
	mu    sync.Mutex //mutual exclusion lock
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++ //critical section
}

func (c *Counter) Value() int {
	return c.value
}
