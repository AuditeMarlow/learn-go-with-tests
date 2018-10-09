package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want int) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		checkArea(t, Rectangle{12, 6}, 72)
	})

	t.Run("circles", func(t *testing.T) {
		checkArea(t, Circle{10}, 314)
	})
}
