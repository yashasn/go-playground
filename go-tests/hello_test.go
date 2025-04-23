package main

import "testing"

/*
t *testing.t is a pointer refers to a struct type from the testing framework. A new instance of this type is created for every test fucntion by the framework.
The framework uses this object to check for status and etc. If * was not there, a copy of the struct type would be used and all the values would be local to the function
and not accessible by the framwork
*/
func TestHello(t *testing.T) {

	// var got string - declaring
	// got = Hello() - initilising
	// := does both at once, infers the type from the value assigned
	got := Hello()
	want := "Hello"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
