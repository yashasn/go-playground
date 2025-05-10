package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
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
}
