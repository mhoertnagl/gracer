package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func NewTuple(x, y, z, w float64) *Tuple {
	return &Tuple{x, y, z, w}
}

func NewVector(x, y, z float64) *Tuple {
	return NewTuple(x, y, z, 0)
}

func NewPoint(x, y, z float64) *Tuple {
	return NewTuple(x, y, z, 1)
}

func (a *Tuple) String() string {
	return fmt.Sprintf("Tuple(%f, %f, %f, %f)", a.x, a.y, a.z, a.w)
}

func (a *Tuple) Add(b *Tuple) *Tuple {
	return NewTuple(a.x+b.x, a.y+b.y, a.z+b.z, a.w+b.w)
}

func (a *Tuple) Sub(b *Tuple) *Tuple {
	return NewTuple(a.x-b.x, a.y-b.y, a.z-b.z, a.w-b.w)
}

func (a *Tuple) Negate() *Tuple {
	return NewTuple(-a.x, -a.y, -a.z, -a.w)
}

func (a *Tuple) Mul(f float64) *Tuple {
	return NewTuple(f*a.x, f*a.y, f*a.z, f*a.w)
}

func (a *Tuple) Div(d float64) *Tuple {
	return NewTuple(a.x/d, a.y/d, a.z/d, a.w/d)
}

func (a *Tuple) Dot(b *Tuple) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w
}

func (a *Tuple) Magnitude() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a *Tuple) Normalize() *Tuple {
	n := a.Magnitude()
	return NewTuple(a.x/n, a.y/n, a.z/n, a.w/n)
}

func (a *Tuple) Cross(b *Tuple) *Tuple {
	return NewVector(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
}

type Color struct {
	r float64
	g float64
	b float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

func (a *Color) String() string {
	return fmt.Sprintf("Color(%f, %f, %f)", a.r, a.g, a.b)
}

func (a *Color) Add(b *Color) *Color {
	return NewColor(a.r+b.r, a.g+b.g, a.b+b.b)
}

func (a *Color) Sub(b *Color) *Color {
	return NewColor(a.r-b.r, a.g-b.g, a.b-b.b)
}

func (a *Color) Scale(f float64) *Color {
	return NewColor(f*a.r, f*a.g, f*a.b)
}

func (a *Color) Mul(b *Color) *Color {
	return NewColor(a.r*b.r, a.g*b.g, a.b*b.b)
}

func (c *Color) RGBA() color.RGBA {
	return color.RGBA{cap(c.r), cap(c.g), cap(c.b), 255}
}

func cap(f float64) uint8 {
	return uint8(math.Min(f, 1) * 255)
}

type Canvas struct {
	m *image.RGBA
}

func NewCanvas(w, h int) *Canvas {
	r := image.Rect(0, 0, w, h)
	m := image.NewRGBA(r)
	return &Canvas{m}
}

func (v *Canvas) Set(x, y int, c *Color) {
	v.m.Set(x, y, c.RGBA())
}

func (v *Canvas) Write(fn string) {
	outFile, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, v.m, &jpeg.Options{Quality: 100})
}

func main() {
	v := NewCanvas(300, 250)
	v.Set(100, 100, NewColor(1, 1, 1))
	v.Write("out.jpg")
	fmt.Println("Done")
}
