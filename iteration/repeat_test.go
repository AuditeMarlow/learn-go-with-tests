package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEquals := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("it can repeat 'a' 5 times", func(t *testing.T) {
		assertEquals(t, Repeat("a", 5), "aaaaa")
	})

	t.Run("it can repeat 'b' 3 times", func(t *testing.T) {
		assertEquals(t, Repeat("b", 3), "bbb")
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
