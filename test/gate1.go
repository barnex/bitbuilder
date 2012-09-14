package main

// Author: Arne Vansteenkiste
// Some code snippets were taken from 
// https://code.google.com/p/appengine-go/

import (
	. "bitbuilder"
)

func main() {
	W, H := 512, 256 // image size
	canvas := NewCanvas(W, H)

	w := 32 // nanowire size
	canvas.SetStroke(w)

	A1 := Point(0, 0+w/2)    // input 1
	B1 := Point(A1.X, H-w/2) // input 2
	A2 := Point(W/3, A1.Y)
	B2 := Point(A2.X, B1.Y)
	canvas.Line(A1, A2)
	canvas.Line(B1, B2)

	C := Point(2*W/3, H/2) // crossing point
	canvas.Line(A2, C)
	canvas.Line(B2, C)

	C2 := Point(W, C.Y) // output
	canvas.Line(C, C2)

	canvas.Encode("gate1.png")
}
