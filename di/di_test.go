package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Niek")

	got := buffer.String()
	want := "Hello, Niek"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
