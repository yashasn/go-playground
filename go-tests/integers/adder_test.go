package integers

/* Since this is non-main package. The name of the package should match teh directory. The directory name should be in lowercase.
https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project */

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(4, 5)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/*
Example functions act as executable documentation. Unlike traditional comments,
these examples are verified by the Go test runner, ensuring they remain accurate as your code evolves.
When you use the go doc tool on your package or function, the content of your example functions
(specifically the part before the // Output: comment and the expected output itself) is included in the generated documentation

Without the //Output this function won't be verified by Go
*/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
