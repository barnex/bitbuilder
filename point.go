package bitbuilder

type Pt struct {
	X, Y int
}

func Point(x, y int) Pt {
	return Pt{x, y}
}
