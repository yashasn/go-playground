package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns faster server", func(t *testing.T) {
		slowServer := makeMockServer(20 * time.Millisecond)
		fastServer := makeMockServer(0)

		// with defer it will now call that function at the end of the containing function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		//cant use these URLs in test. Need to mock them
		// slowURL := "http://www.facebook.com"
		// fastURL := "http://www.quii.dev"

		want := fastURL
		got, _ := RacerOptimised(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns timeout", func(t *testing.T) {
		slowServer := makeMockServer(20 * time.Millisecond)
		fastServer := makeMockServer(20 * time.Millisecond)

		// with defer it will now call that function at the end of the containing function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		//cant use these URLs in test. Need to mock them
		// slowURL := "http://www.facebook.com"
		// fastURL := "http://www.quii.dev"

		_, err := ConfigurableRacer(slowURL, fastURL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeMockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
