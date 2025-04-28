package main

import (
	"testing"
)

/*
t *testing.t is a pointer refers to a struct type from the testing framework. A new instance of this type is created for every test fucntion by the framework.
The framework uses this object to check for status and etc. If * was not there, a copy of the struct type would be used and all the values would be local to the function
and not accessible by the framwork
*/

// The test function must start with the word Test
func DummyHello(t *testing.T) {

	// var got string - declaring
	// got = Hello() - initialising
	// := does both at once, infers the type from the value assigned
	recp := "Yashas"
	got := Hello(recp, "")
	want := "Hello, " + recp

	if got != want {
		t.Errorf("got %q want %q", got, want)
		//%q prints string with "". Ex: got=" hello", but %s prints got= hello. With %q easy to find whitespace and special characters
	}
}

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		recp := "Yashas"
		got := Hello(recp, "")
		want := "Hello, " + recp

		assertMessage(t, got, want)
	})
	t.Run("Saying Hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
	t.Run("Saying Hello in Kannada", func(t *testing.T) {
		got := Hello("Yash", "Kannada")
		want := "Namaskara, Yash"

		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
