package main

import (
	"reflect"
	"testing"
)

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

func TestSumAllTails(t *testing.T) {
	//functions can be assigned to variables !!
	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		//NOTE : can't use == for comparing slices
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("Sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 1, 1})
		want := []int{5, 2}

		checkSum(t, got, want)
	})
	t.Run("Sum of slices with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 1, 1})
		want := []int{0, 2}
		checkSum(t, got, want)
	})
}
