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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
