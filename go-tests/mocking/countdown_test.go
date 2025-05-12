package main

import (
	"reflect"
	"testing"
)

type SpyOperations struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

// These mock the order of operations. This helps us test the Countdown func more finely
func (s *SpyOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {

	t.Run("countdown", func(t *testing.T) {
		spyOps := &SpyOperations{}
		Countdown(spyOps, spyOps)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spyOps.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyOps.Calls)
		}
	})
}
