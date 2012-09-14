package main

// Author: Arne Vansteenkiste

import (
	. "github.com/barnex/bitbuilder"
)

func main() {
	W, H := 512, 256 // image size
	canvas := NewCanvas(W, H)

	w := 32 // nanowire width
	canvas.SetStroke(w, Round)

	A1 := Point{w, 0 + w/2}    // input 1
	B1 := Point{A1.X, H - w/2} // input 2
	A2 := Point{W / 3, A1.Y}
	B2 := Point{A2.X, B1.Y}
	canvas.Line(A1, A2)
	canvas.Line(B1, B2)

	C := Point{2 * W / 3, H / 2} // crossing point
	canvas.Line(A2, C)
	canvas.Line(B2, C)

	C2 := Point{W - w, C.Y} // output
	canvas.SetStroke(2*w, Butt)
	canvas.Line(C, C2)

	// triangle at the output
	a := C2.Add(0, w)
	b := C2.Add(w, 0)
	c := C2.Add(0, -w)
	triangle := Poly{a, b, c}
	canvas.Fill(triangle)

	canvas.Encode("gate2.png")
}
