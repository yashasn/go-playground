package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		/*
			The channel can hold one string value without requiring a receiver to be ready.
			Once one value is in the channel, any further sends (data <- "value") will block until the value is received.
			Example :
			data := make(chan string, 1)

			data <- "hello" // success, buffer has space

			data <- "world" // blocks here, buffer is full and no receiver
		*/

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}

	}
}
