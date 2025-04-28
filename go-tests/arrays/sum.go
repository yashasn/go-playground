package main

/*
NOTE : for Arrays size is encoded in its type i.e [5]int.
If we try to pass an [4]int into a function that expects [5]int, it won't compile.
*/
func SumWArrays(numbers [5]int) int {
	sum := 0
	//range is used to iterate over an array. Returns index and the value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
Slices don't have fixed size. Slice is implemented as a lightweight struct (Data- ptr to data, Len, Cap)
Passing a slice to a function copies the slice header, but not the actual data
*/
func SumWSlices(numbers []int) int {
	sum := 0
	//range is used to iterate over an array. Returns index and the value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersSlice ...[]int) []int {
	var sums []int
	for _, nums := range numbersSlice {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tails := nums[1:]
			sums = append(sums, SumWSlices(tails))
		}

	}
	return sums
}
