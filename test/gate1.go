package main

// Author: Arne Vansteenkiste
// Some code snippets were taken from 
// https://code.google.com/p/appengine-go/

import (
	"bitbuilder"
)

func main() {
	canvas := bitbuilder.NewCanvas(512, 256)
	canvas.SetStroke(16)
	canvas.Line(64, 32, 256, 128)
	canvas.Encode("gate1.png")
}

