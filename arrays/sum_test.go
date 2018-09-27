package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertEquals := func(t *testing.T, got, want int, numbers []int) {
		if got != want {
			t.Errorf("got '%d', want '%d', given '%v'", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		assertEquals(t, Sum(numbers), 15, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		assertEquals(t, Sum(numbers), 6, numbers)
	})

	t.Run("sum many collections' trails", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v', want '%v'", got, want)
		}
	})

	t.Run("safely sum up empty slices", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{8, 2})
		want := []int{0, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v', want '%v'", got, want)
		}
	})
}
