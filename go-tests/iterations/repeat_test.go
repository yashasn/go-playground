package iterations

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Unoptimised Repeat function", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Optimised Repeat function", func(t *testing.T) {
		got := RepeatOptimised("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

/*
benchmarks are special types of tests that measure how fast or efficient your code is.
b.N is automatically controlled by the Go test framework.
Go runs the function repeatedly and adjusts b.N until it can measure performance reliably.
By default benchmarks are run sequentially.
cmd to run:- go test -bench=. (Runs all Benchmark function, can also specify Benchmark fucntion)
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func BenchmarkRepeatOptimised(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatOptimised("a", 5)
	}
}
