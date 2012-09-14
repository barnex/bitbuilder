package main

// Author: Arne Vansteenkiste

import (
	. "github.com/barnex/bitbuilder"
)

func main() {
	W, H := 512, 512 // image size
	canvas := NewCanvas(W, H)

	w := 32 // nanowire width

	A1 := Point{0, 0 + w/2}    // input 1
	B1 := Point{A1.X, H - w/2} // input 2
	A2 := Point{W / 3, A1.Y}
	B2 := Point{A2.X, B1.Y}
	C := Point{2 * W / 3, H / 2} // crossing point
	C2 := Point{W, C.Y}          // output

	var wire1 Wire
	wire1.Add(A1, w)
	wire1.Add(A2, w)
	wire1.Add(C, w/2)
	wire1.Add(C2, w/2)
	wire1.Draw(canvas)

	var wire2 Wire
	wire2.Add(B1, w)
	wire2.Add(B2, w)
	wire2.Add(C, w/2)
	wire2.Draw(canvas)

	canvas.Encode("wire.png")
}
