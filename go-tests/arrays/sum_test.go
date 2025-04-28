package main

import "testing"

/* Code coverage
go test -cover
*/

func TestSum(t *testing.T) {
	t.Run("Fixed size collections with Arrays", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := SumWArrays(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
	t.Run("Fixed size collections with Arrays", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := SumWSlices(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}
