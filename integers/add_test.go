package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertEquals := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("it can add 2 and 2 and expect 4", func(t *testing.T) {
		assertEquals(t, Add(2, 2), 4)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
