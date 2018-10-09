package shapes

import "math"

type Shape interface {
	Area() int
}

type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius int
}

type Triangle struct {
	Base   int
	Height int
}

func Perimeter(rectangle Rectangle) int {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (c Circle) Area() int {
	return int(math.Pi * float64(c.Radius*c.Radius))
}

func (t Triangle) Area() int {
	return int(float64(t.Base*t.Height) * 0.5)
}
