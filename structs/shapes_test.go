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
		{shape: Rectangle{Width: 12, Height: 6}, want: 72},
		{shape: Circle{Radius: 10}, want: 314},
		{shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
