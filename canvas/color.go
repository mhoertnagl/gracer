package canvas

import (
	"fmt"
	"image/color"
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Color alg.Vector

func NewColor(r, g, b float64) Color {
	return Color(alg.NewVector(r, g, b))
}

func Black() Color {
	return NewColor(0, 0, 0)
}

func (c Color) String() string {
	return fmt.Sprintf("Color(%f, %f, %f)", c[0], c[1], c[2])
}

func (a Color) Add(b Color) Color {
	va := alg.Vector(a)
	vb := alg.Vector(b)
	return Color(va.Add(vb))
}

func (a Color) Sub(b Color) Color {
	va := alg.Vector(a)
	vb := alg.Vector(b)
	return Color(va.Sub(vb))
}

func (c Color) Scale(f float64) Color {
	vc := alg.Vector(c)
	return Color(vc.Mult(f))
}

func (a Color) Mul(b Color) Color {
	va := alg.Vector(a)
	vb := alg.Vector(b)
	return Color(va.Had(vb))
}

func (c Color) RGBA() color.RGBA {
	return color.RGBA{cap(c[0]), cap(c[1]), cap(c[2]), 255}
}

func cap(f float64) uint8 {
	return uint8(math.Min(f, 1) * 255)
}
