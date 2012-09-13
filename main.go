package main

// Author: Arne Vansteenkiste
// Some code snippets were taken from 
// https://code.google.com/p/appengine-go/

import (
	"code.google.com/p/freetype-go/freetype/raster"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	w, h := 800, 600
	img := image.NewRGBA(image.Rect(0, 0, h, w))
	p := raster.NewRGBAPainter(img)
	p.SetColor(color.RGBA{0, 0, 0, 255})
	r := raster.NewRasterizer(w, h)
	r.UseNonZeroWinding = true

	width := fix(40)
	var path raster.Path
	path.Start(pt(0, 0))
	path.Add1(pt(100, 100))
	path.Add1(pt(200, 100))
	raster.Stroke(r, path, width, nil, nil)

	r.Rasterize(p)
	encode(img, "out.png")
}

func encode(img image.Image, fname string) {
	out, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	check(err)
	defer out.Close()
	err = png.Encode(out, img)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func pt(x, y int) raster.Point {
	return raster.Point{fix(x), fix(y)}
}

func fix(x int) raster.Fix32 {
	return raster.Fix32(x << 8)
}
