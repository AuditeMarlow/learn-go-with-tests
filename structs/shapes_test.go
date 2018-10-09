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
		name  string
		shape Shape
		want  int
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %v, want %v", tt.shape, got, tt.want)
			}
		})
	}
}
