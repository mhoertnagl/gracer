package main

import "fmt"

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

func main() {
	fmt.Println("0")
}
