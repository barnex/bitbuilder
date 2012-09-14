package bitbuilder

// A point in 2D space.
type Point struct {
	X, Y int
}

// Returns a new point translated by dx, dy.
func (p Point) Add(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}
