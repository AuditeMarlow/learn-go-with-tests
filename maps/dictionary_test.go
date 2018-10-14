package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test."}

	got := Search(dictionary, "test")
	want := "This is just a test."

	assertEquals(t, got, want)
}

func assertEquals(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
