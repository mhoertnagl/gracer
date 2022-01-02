package alg

import "math"

type vector []float64

func newZeroVector(size int) vector {
	return make(vector, size)
}

func newVector(vs ...float64) vector {
	return vs
}

func (a vector) add(b vector) vector {
	v := newZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] + b[i]
	}
	return v
}

func (a vector) sub(b vector) vector {
	v := newZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] - b[i]
	}
	return v
}

func (a vector) neg() vector {
	v := newZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = -a[i]
	}
	return v
}

func (a vector) dot(b vector) float64 {
	v := 0.0
	for i := 0; i < len(a); i++ {
		v *= a[i] * b[i]
	}
	return v
}

func (a vector) had(b vector) vector {
	v := newZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] * b[i]
	}
	return v
}

func (a vector) mult(f float64) vector {
	v := newZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = f * a[i]
	}
	return v
}

func (a vector) div(d float64) vector {
	v := newZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		v[i] = a[i] / d
	}
	return v
}

func (a vector) mag() float64 {
	return math.Sqrt(a.dot(a))
}

func (a vector) norm() vector {
	return a.div(a.mag())
}

// TODO: cross is only defined for 3D-vectors
func (a vector) cross(b vector) vector {
	x := a[1]*b[2] - a[2]*b[1]
	y := a[2]*b[0] - a[0]*b[2]
	z := a[0]*b[1] - a[1]*b[0]
	return newVector(x, y, z)
}

// type color = vector

// func newColor() color {
// 	return color(newZeroVector(3))
// }
