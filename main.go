package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"strings"
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

func (c *Color) String() string {
	return fmt.Sprintf("Color(%f, %f, %f)", c.r, c.g, c.b)
}

func (a *Color) Add(b *Color) *Color {
	return NewColor(a.r+b.r, a.g+b.g, a.b+b.b)
}

func (a *Color) Sub(b *Color) *Color {
	return NewColor(a.r-b.r, a.g-b.g, a.b-b.b)
}

func (c *Color) Scale(f float64) *Color {
	return NewColor(f*c.r, f*c.g, f*c.b)
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

type Matrix struct {
	m [][]float64
}

var Id4 = NewMatrix(
	Row(1, 0, 0, 0),
	Row(0, 1, 0, 0),
	Row(0, 0, 1, 0),
	Row(0, 0, 0, 1),
)

func Row(rs ...float64) []float64 {
	return rs
}

func NewMatrix(m ...[]float64) *Matrix {
	sz := len(m)
	for _, row := range m {
		if len(row) != sz {
			panic(fmt.Sprintf("Matrix [%d %d] is not square", sz, len(row)))
		}
	}
	return &Matrix{m}
}

func NewEmptyMatrix(size int) *Matrix {
	m := make([][]float64, size)
	for i := range m {
		m[i] = make([]float64, size)
	}
	return &Matrix{m}
}

func (m *Matrix) String() string {
	var b strings.Builder
	sz := len(m.m)
	b.WriteString("Matrix(")
	for r := 0; r < sz; r++ {
		b.WriteString("\n  ")
		for c := 0; c < sz; c++ {
			v := m.m[r][c]
			b.WriteString(fmt.Sprintf("%8.6f  ", v))
		}
	}
	b.WriteString("\n)")
	return b.String()
}

func (a *Matrix) MatMul(b *Matrix) *Matrix {
	sza := len(a.m)
	szb := len(b.m)
	if sza != szb {
		panic(fmt.Sprintf("Matrix sizes mismatch [%d] vs. [%d]", sza, szb))
	}
	m := NewEmptyMatrix(sza)
	for r := 0; r < sza; r++ {
		for c := 0; c < sza; c++ {
			for l := 0; l < sza; l++ {
				m.m[r][c] += a.m[r][l] * b.m[l][c]
			}
		}
	}
	return m
}

func (a *Matrix) TupMul(b *Tuple) *Tuple {
	sza := len(a.m)
	if sza != 4 {
		panic(fmt.Sprintf("Matrix sizes mismatch [%d] vs. [%d]", sza, 4))
	}
	x := a.m[0][0]*b.x + a.m[0][1]*b.y + a.m[0][2]*b.z + a.m[0][3]*b.w
	y := a.m[1][0]*b.x + a.m[1][1]*b.y + a.m[1][2]*b.z + a.m[1][3]*b.w
	z := a.m[2][0]*b.x + a.m[2][1]*b.y + a.m[2][2]*b.z + a.m[2][3]*b.w
	w := a.m[3][0]*b.x + a.m[3][1]*b.y + a.m[3][2]*b.z + a.m[3][3]*b.w
	return NewTuple(x, y, z, w)
}

func (a *Matrix) Transpose() *Matrix {
	sz := len(a.m)
	m := NewEmptyMatrix(sz)
	for r := 0; r < sz; r++ {
		for c := 0; c < sz; c++ {
			m.m[c][r] = a.m[r][c]
		}
	}
	return m
}

func (m *Matrix) LUFactorize() (*Matrix, *Matrix) {
	sz := len(m.m)
	l := NewEmptyMatrix(sz)
	u := NewEmptyMatrix(sz)
	for i := 0; i < sz; i++ {
		for j := i; j < sz; j++ {
			d := 0.0
			e := 0.0
			for k := 0; k < i; k++ {
				d += l.m[j][k] * u.m[k][i]
				e += l.m[i][k] * u.m[k][j]
			}
			l.m[j][i] = m.m[j][i] - d
			u.m[i][j] = (m.m[i][j] - e) / l.m[i][i]
		}
	}
	return l, u
}

// func lu(m [][]float64) ([][]float64, [][]float64) {
// 	sz := len(m)
// 	l := emptyArray2D(sz)
// 	u := emptyArray2D(sz)
// 	for i := 0; i < sz; i++ {
// 		for j := i; j < sz; j++ {
// 			d := 0.0
// 			e := 0.0
// 			for k := 0; k < i; k++ {
// 				d += l.m[j][k] * u.m[k][i]
// 				e += l.m[i][k] * u.m[k][j]
// 			}
// 			l[j][i] = m[j][i] - d
// 			u[i][j] = (m[i][j] - e) / l[i][i]
// 		}
// 	}
// 	return l, u
// }

// func emptyArray2D(size int) [][]float64 {
// 	m := make([][]float64, size)
// 	for i := range m {
// 		m[i] = make([]float64, size)
// 	}
// 	return m
// }

func main() {
	v := NewCanvas(300, 250)
	v.Set(100, 100, NewColor(1, 1, 1))
	v.Write("out.jpg")
	fmt.Println("Done")
}
