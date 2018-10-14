package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	wantedCalls := 3

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s', wanted '%s'", got, want)
	}

	if sleeper.Calls != wantedCalls {
		t.Errorf("got %d calls to sleeper, wanted %d", sleeper.Calls, wantedCalls)
	}
}
