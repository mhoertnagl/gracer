package alg

import (
	"fmt"
	"math"
	"strings"
)

type Vector []float64

var Origin = NewPoint(0, 0, 0)

// TODO: X, Y, Z, NX, NY, NZ

func NewZeroVector(size int) Vector {
	return make(Vector, size)
}

// TODO: NewVector should be Vector3 and Vector should be renamed.
func NewVector(vs ...float64) Vector {
	return vs
}

func NewVector3(x, y, z float64) Vector {
	return NewVector(x, y, z, 0)
}

func NewPoint(x, y, z float64) Vector {
	return NewVector(x, y, z, 1)
}

func (a Vector) String() string {
	var b strings.Builder
	b.WriteString("Vector(")
	for i := 0; i < len(a); i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%8.6f", a[i]))
	}
	b.WriteString(")")
	return b.String()
}

func (a Vector) Add(b Vector) Vector {
	v := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] + b[i]
	}
	return v
}

func (a Vector) Sub(b Vector) Vector {
	v := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] - b[i]
	}
	return v
}

func (a Vector) Neg() Vector {
	v := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = -a[i]
	}
	return v
}

func (a Vector) Mult(f float64) Vector {
	v := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = f * a[i]
	}
	return v
}

func (a Vector) Div(d float64) Vector {
	v := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] / d
	}
	return v
}

func (a Vector) Dot(b Vector) float64 {
	v := 0.0
	for i := 0; i < len(a); i++ {
		v += a[i] * b[i]
	}
	return v
}

// TODO: cross is only defined for 3D-Vectors
func (a Vector) Cross(b Vector) Vector {
	x := a[1]*b[2] - a[2]*b[1]
	y := a[2]*b[0] - a[0]*b[2]
	z := a[0]*b[1] - a[1]*b[0]
	return NewVector3(x, y, z)
}

func (a Vector) Had(b Vector) Vector {
	v := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] * b[i]
	}
	return v
}

func (a Vector) Mag() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Norm() Vector {
	return a.Div(a.Mag())
}

func (a Vector) Reflect(n Vector) Vector {
	return a.Sub(n.Mult(2 * a.Dot(n)))
}
