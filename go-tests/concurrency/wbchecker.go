package concurrency

/*
Goroutine runs a seperate process. Only way to start a goroutine is to put go in front of a function call.
We often use anonymous functions when we want to start a goroutine.
Anonymous funcs can be immediately invoked from where they are declared.
Example:
go func(msg string) {
    fmt.Println(msg)
}("Hello from goroutine")
*/

/*
The below implmentation leads to race condition (go test -race, to findout more) because mutiple goroutines write to same shared variable.
Go maps are not safe for concurrent writes. Even if multiple goroutines write to different keys,
the internal structure of the map may still be modified (e.g., during resizing, rehashing, or key indexing),
and that leads to data races and undefined behavior.

We can avoid this using Channels.
*/

type WebsiteChecker func(string) bool

type result struct {
	url    string
	status bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			//send value to channel
			resultChannel <- result{url, wc(url)}
		}()
	}
	for i := 0; i < len(urls); i++ {
		//receive results from the channel one by one
		//perform the map write synchronously, in a single thread — no race condition!
		//this is a blocking call, as it is waiting for a value in the channel
		r := <-resultChannel
		results[r.url] = r.status
	}
	return results
}
