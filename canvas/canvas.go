package canvas

import (
	"image"
	"image/jpeg"
	"os"
)

type Canvas struct {
	m *image.RGBA
}

func NewCanvas(w, h int) *Canvas {
	r := image.Rect(0, 0, w, h)
	m := image.NewRGBA(r)
	return &Canvas{m}
}

func (v *Canvas) Set(x, y int, c Color) {
	v.m.Set(x, y, c.RGBA())
}

func (v *Canvas) WriteToFile(name string) {
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, v.m, &jpeg.Options{Quality: 100})
}
