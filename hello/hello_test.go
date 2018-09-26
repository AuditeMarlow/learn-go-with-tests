package main

import "testing"

func TestHello(t *testing.T) {
	assertEquals := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		assertEquals(t, Hello("Niek", ""), "Hello, Niek")
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		assertEquals(t, Hello("", ""), "Hello, World")
	})

	t.Run("in Spanish!", func(t *testing.T) {
		assertEquals(t, Hello("Niek", "Spanish"), "Hola, Niek")
	})

	t.Run("in French!", func(t *testing.T) {
		assertEquals(t, Hello("Niek", "French"), "Bonjour, Niek")
	})
}
