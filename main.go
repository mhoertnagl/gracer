package main

import (
	"fmt"
	"math"
)

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (a *Tuple) String() string {
	return fmt.Sprintf("Tuple(%f, %f, %f, %f)", a.x, a.y, a.z, a.w)
}

func tuple(x, y, z, w float64) *Tuple {
	return &Tuple{x, y, z, w}
}

func vector(x, y, z float64) *Tuple {
	return tuple(x, y, z, 0)
}

func point(x, y, z float64) *Tuple {
	return tuple(x, y, z, 1)
}

func add(a, b *Tuple) *Tuple {
	return tuple(a.x+b.x, a.y+b.y, a.z+b.z, a.w+b.w)
}

func sub(a, b *Tuple) *Tuple {
	return tuple(a.x-b.x, a.y-b.y, a.z-b.z, a.w-b.w)
}

func negate(a *Tuple) *Tuple {
	return tuple(-a.x, -a.y, -a.z, -a.w)
}

func mul(f float64, a *Tuple) *Tuple {
	return tuple(f*a.x, f*a.y, f*a.z, f*a.w)
}

func div(a *Tuple, d float64) *Tuple {
	return tuple(a.x/d, a.y/d, a.z/d, a.w/d)
}

func dot(a, b *Tuple) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w
}

func magnitude(a *Tuple) float64 {
	return math.Sqrt(dot(a, a))
}

func normalize(a *Tuple) *Tuple {
	n := magnitude(a)
	return tuple(a.x/n, a.y/n, a.z/n, a.w/n)
}

func cross(a, b *Tuple) *Tuple {
	return vector(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
}

// func NewTuple(x, y, z, w float64) *Tuple {
// 	return &Tuple{x, y, z, w}
// }

// func NewVector(x, y, z float64) *Tuple {
// 	return NewTuple(x, y, z, 0)
// }

// func NewPoint(x, y, z float64) *Tuple {
// 	return NewTuple(x, y, z, 1)
// }

// func (a *Tuple) String() string {
// 	return fmt.Sprintf("Tuple(%f, %f, %f, %f)", a.x, a.y, a.z, a.w)
// }

// func (a *Tuple) Add(b *Tuple) *Tuple {
// 	return NewTuple(a.x+b.x, a.y+b.y, a.z+b.z, a.w+b.w)
// }

// func (a *Tuple) Sub(b *Tuple) *Tuple {
// 	return NewTuple(a.x-b.x, a.y-b.y, a.z-b.z, a.w-b.w)
// }

// func (a *Tuple) Negate() *Tuple {
// 	return NewTuple(-a.x, -a.y, -a.z, -a.w)
// }

// func (a *Tuple) Mul(f float64) *Tuple {
// 	return NewTuple(f*a.x, f*a.y, f*a.z, f*a.w)
// }

// func (a *Tuple) Div(d float64) *Tuple {
// 	return NewTuple(a.x/d, a.y/d, a.z/d, a.w/d)
// }

// func (a *Tuple) Dot(b *Tuple) float64 {
// 	return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w
// }

// func (a *Tuple) Magnitude() float64 {
// 	return math.Sqrt(a.Dot(a))
// }

// func (a *Tuple) Normalize() *Tuple {
// 	n := a.Magnitude()
// 	return NewTuple(a.x/n, a.y/n, a.z/n, a.w/n)
// }

// func (a *Tuple) Cross(b *Tuple) *Tuple {
// 	return NewVector(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
// }

func main() {
	fmt.Println("0")
}
