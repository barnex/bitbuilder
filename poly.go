package bitbuilder

// Arbitrary polygon defined by array of points.
type Poly []Point

// Adds a point to the polygon.
func (p *Poly) Add(x ...Point) {
	*p = append(*p, x...)
}

// Resets the polygon,
// removing all points.
func (p *Poly) Reset() {
	*p = (*p)[:0]
}

// Draws the polygon (filled).
func (c *Canvas) Fill(p Poly) {
	c.resetPath()
	c.path.Start(pt(p[0]))
	for i := 1; i < len(p); i++ {
		c.path.Add1(pt(p[i]))
	}
	c.path.Add1(pt(p[0]))
	c.fillPath()
	c.resetPath()
}
