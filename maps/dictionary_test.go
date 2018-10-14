package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test."}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test."

		assertEquals(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("bogus")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	var definition = "This is just a test."

	dictionary := Dictionary{}
	dictionary.Add("test", definition)

	got, err := dictionary.Search("test")
	want := definition

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertEquals(t, got, want)
}

func assertEquals(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
