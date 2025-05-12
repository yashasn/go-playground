package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(10 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {

	t.Run("positive case", func(t *testing.T) {
		//why &SpyStore and not just SpyStore, because Fetch() have a pointer receiver and not value receiver. So only *SpyStore is considered to implement Fetch
		//So only *SpyStore implments the interface and not SpyStore
		data := "Hello, World"
		spyStore := &SpyStore{response: data}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("store should cancel work if req is cancelled", func(t *testing.T) {

		data := "Hello, World"
		spyStore := &SpyStore{response: data}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !spyStore.cancelled {
			t.Error("store was not told to cancel")
		}

	})
}
