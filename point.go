package bitbuilder

import "math"

// A point in 2D space.
type Point struct {
	X, Y int
}

// Returns a new point translated by dx, dy.
func (p Point) Add2(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}
// Returns a + b.
func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

// Returns a - b.
func(a Point)Sub(b Point)Point{
	return Point{a.X-b.X, a.Y-b.Y}
}

// Re-scales a's length (from origin).
func(a Point)Rescale(length float64)Point{
	x, y := float64(a.X), float64(a.Y)
	n := length / math.Sqrt(x*x + y*y)
	x *= n
	y *= n
	return Point{int(x), int(y)}
}

func(a Point) Rot90()Point{
	return Point{a.Y, -a.X}
}

func(a Point) Rot270()Point{
	return Point{-a.Y, a.X}
}
