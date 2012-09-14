package bitbuilder

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

type Canvas struct {
	*image.RGBA
	painter     *raster.RGBAPainter
	rasterizer  *raster.Rasterizer
	strokewidth raster.Fix32
	strokecap   raster.Capper
	path        raster.Path
}

func NewCanvas(w, h int) *Canvas {
	c := new(Canvas)
	c.RGBA = image.NewRGBA(image.Rect(0, 0, w, h))
	c.painter = raster.NewRGBAPainter(c.RGBA)
	c.rasterizer = raster.NewRasterizer(w, h)
	c.rasterizer.UseNonZeroWinding = true
	c.SetColor(color.Black)
	c.path = make(raster.Path, 0, 100)
	c.resetPath()
	c.SetStroke(1, Round)
	return c
}

func (c *Canvas) SetColor(col color.Color) {
	c.painter.SetColor(col)
}

var (
	Round  = raster.RoundCapper
	Square = raster.SquareCapper
	Butt   = raster.ButtCapper
)

func (c *Canvas) SetStroke(width int, cap_ raster.Capper) {
	c.strokewidth = fix32(width)
	c.strokecap = cap_
}

func (c *Canvas) Line(a, b Pt) {
	c.path.Start(pt(a.X, a.Y))
	c.path.Add1(pt(b.X, b.Y))
	c.strokePath()
	c.resetPath()
}

func (c *Canvas) resetPath() {
	c.path = c.path[:0]
}

func (c *Canvas) strokePath() {
	raster.Stroke(c.rasterizer, c.path, c.strokewidth, c.strokecap, nil)
	c.rasterizer.Rasterize(c.painter)
}

func (c *Canvas) Encode(fname string) {
	out, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	check("topng", err)
	defer out.Close()
	err = png.Encode(out, c.RGBA)
	check("topng", err)
}

func check(msg string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg, ":", err)
		os.Exit(1)
	}
}

func pt(x, y int) raster.Point {
	return raster.Point{fix32(x), fix32(y)}
}

func fix32(x int) raster.Fix32 {
	return raster.Fix32(x << 8)
}
