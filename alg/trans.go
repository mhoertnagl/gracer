package alg

import "math"

func Translation(x, y, z float64) Matrix {
	return NewMatrix(
		Row(1, 0, 0, x),
		Row(0, 1, 0, y),
		Row(0, 0, 1, z),
		Row(0, 0, 0, 1),
	)
}

func Scaling(x, y, z float64) Matrix {
	return NewMatrix(
		Row(x, 0, 0, 0),
		Row(0, y, 0, 0),
		Row(0, 0, z, 0),
		Row(0, 0, 0, 1),
	)
}

func RotationX(r float64) Matrix {
	c := math.Cos(r)
	s := math.Sin(r)
	return NewMatrix(
		Row(1, 0, 0, 0),
		Row(0, c, -s, 0),
		Row(0, s, c, 0),
		Row(0, 0, 0, 1),
	)
}

func RotationY(r float64) Matrix {
	c := math.Cos(r)
	s := math.Sin(r)
	return NewMatrix(
		Row(c, 0, s, 0),
		Row(0, 1, 0, 0),
		Row(-s, 0, c, 0),
		Row(0, 0, 0, 1),
	)
}

func RotationZ(r float64) Matrix {
	c := math.Cos(r)
	s := math.Sin(r)
	return NewMatrix(
		Row(c, -s, 0, 0),
		Row(s, c, 0, 0),
		Row(0, 0, 1, 0),
		Row(0, 0, 0, 1),
	)
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return NewMatrix(
		Row(1, xy, xz, 0),
		Row(yx, 1, yz, 0),
		Row(zx, zy, 1, 0),
		Row(0, 0, 0, 1),
	)
}

func (a Matrix) Translate(x, y, z float64) Matrix {
	return Translation(x, y, z).MultMat(a)
}

func (a Matrix) Scale(x, y, z float64) Matrix {
	return Scaling(x, y, z).MultMat(a)
}

func (a Matrix) RotateX(r float64) Matrix {
	return RotationX(r).MultMat(a)
}

func (a Matrix) RotateY(r float64) Matrix {
	return RotationY(r).MultMat(a)
}

func (a Matrix) RotateZ(r float64) Matrix {
	return RotationZ(r).MultMat(a)
}

func (a Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Shearing(xy, xz, yx, yz, zx, zy).MultMat(a)
}
