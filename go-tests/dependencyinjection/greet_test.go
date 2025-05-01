package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	GreetDI(&buffer, "Yash")

	got := buffer.String()
	want := "Hello,Yash"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
