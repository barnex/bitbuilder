package main

import(
	 "code.google.com/p/freetype-go/freetype/raster"
	 "image"
	 "image/png"
	 "image/color"
	 "fmt"
	 "os"
)

func main(){
	w, h := 200, 100
	img := image.NewRGBA(image.Rect(0, 0, h, w))
	p := raster.NewRGBAPainter(img)
 	p.SetColor(color.RGBA{0, 0, 0, 255})
	r := raster.NewRasterizer(w, h)
	r.Rasterize(p)
	encode(img, "out.png")
}

func encode(img image.Image, fname string){
	out, err := os.OpenFile(fname, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0755)
	check(err)
	defer out.Close()
	err = png.Encode(out, img)
	check(err)
}

func check(err error){
	if err != nil{
		fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
	}
}
