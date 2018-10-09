package shapes

import "math"

type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius int
}

func Perimeter(rectangle Rectangle) int {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (c Circle) Area() int {
	return int(math.Pi * float64(c.Radius) * float64(c.Radius))
}
