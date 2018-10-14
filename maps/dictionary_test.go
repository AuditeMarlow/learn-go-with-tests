package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test."}

	t.Run("known word", func(t *testing.T) {
		got, _ := Search(dictionary, "test")
		want := "This is just a test."

		assertEquals(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := Search(dictionary, "bogus")
		want := "Could not find the word you were looking for."

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertEquals(t, err.Error(), want)
	})
}

func assertEquals(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
