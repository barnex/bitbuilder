package bitbuilder

import(
	"image/color"
)
type Wire struct {
	poly  Poly
	width []int
}

func (w *Wire) Add(p Point, width int) {
	w.poly.Add(p)
	w.width = append(w.width, width)
}

func (w *Wire) Draw(c *Canvas) {
	wire := w.poly
	phase := 0.
	stride := 30. // TODO

	c.SetColor(color.RGBA{0, 0, 0, 255})
//	for i := 0; i < len(wire)-1; i++ {
//		c.WireSeg(wire[i], wire[i+1], w.width[i], w.width[i+1])
//		if i > 0 {
//			c.Circle(wire[i], w.width[i])
//		}
//	}

	c.SetColor(color.RGBA{255, 0, 0, 255})
	for i := 0; i < len(wire)-1; i++ {
		phase = c.Notches(wire[i], wire[i+1], w.width[i], w.width[i+1], stride, phase)
	}
	c.SetColor(color.RGBA{0, 0, 0, 255})
}

func(c*Canvas)Notches(a, b Point, w1, w2 int, stride, phase float64) (newphase float64){
	d1 := b.Sub(a).Normalize(float64(w1) / 2)
	d2 := b.Sub(a).Normalize(float64(w2) / 2)
	x1 := a.Add(d1.Rot90())
	x2 := b.Add(d2.Rot90())
	x3 := b.Add(d2.Rot270())
	x4 := a.Add(d1.Rot270())
	l := b.Sub(a).Length()
	
	dir1 := x2.Sub(x1).Normalize(stride)
	dir2 := x3.Sub(x4).Normalize(stride)
	n1 := x1
	n2 := x4
	for t := -phase; t < l / stride; t ++{
		c.Circle(n1, 5)	
		c.Circle(n2, 5)	
		n1 = n1.Add(dir1)
		n2 = n2.Add(dir2)
	}
	phase += l/stride
	for phase > 1{phase--}
	return 0// todo
}

// Draws a wire segment between a and b,
// with a linearly varying thickness.
func (c *Canvas) WireSeg(a, b Point, w1, w2 int) {
	d1 := b.Sub(a).Normalize(float64(w1) / 2)
	d2 := b.Sub(a).Normalize(float64(w2) / 2)
	x1 := a.Add(d1.Rot90())
	x2 := b.Add(d2.Rot90())
	x3 := b.Add(d2.Rot270())
	x4 := a.Add(d1.Rot270())
	c.Fill(Poly{x1, x2, x3, x4})
}

func (c *Canvas) Circle(center Point, r int) {
	c.SetStroke(r, Round)
	c.resetPath()
	c.path.Start(pt(center))
	c.path.Add1(pt(center.Add2(1, 0)))
	c.strokePath()
	c.resetPath()
}
