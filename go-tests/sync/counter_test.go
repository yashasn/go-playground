package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	wantedCount := 1000
	counter := Counter{}

	// Wait groups are used to synchronise concurrent processes
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	got := counter.Value()
	if got != wantedCount {
		t.Errorf("got %d want %d", got, wantedCount)
	}
}
