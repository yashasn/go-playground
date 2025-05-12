package iterations

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated

	/* strings in golang are IMMUTABLE! So every concatenation involves copying memory and created new string. Impacts performance.
	Use StringBuilder instead minimizes memory.
	*/
}

func RepeatOptimised(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
