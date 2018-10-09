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
	areaTests := []struct {
		shape Shape
		want  int
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 314},
		{Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
