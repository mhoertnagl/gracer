package alg

import "testing"

func TestCalculateInverse1(t *testing.T) {
	e := NewMatrix(
		Row(-0.15385, -0.15385, -0.28205, -0.53846),
		Row(-0.07692, 0.12308, 0.02564, 0.03077),
		Row(0.35897, 0.35897, 0.43590, 0.92308),
		Row(-0.69231, -0.69231, -0.76923, -1.92308),
	)
	x := NewMatrix(
		Row(8, -5, 9, 2),
		Row(7, 5, 6, 1),
		Row(-6, 0, 9, 6),
		Row(-3, 0, -9, -4),
	)
	a := Inverse(x)
	AssertMatrixEqual(t, a, e)
}

func TestCalculateInverse2(t *testing.T) {
	e := NewMatrix(
		Row(-0.04074, -0.07778, 0.14444, -0.22222),
		Row(-0.07778, 0.03333, 0.36667, -0.33333),
		Row(-0.02901, -0.14630, -0.10926, 0.12963),
		Row(0.17778, 0.06667, -0.26667, 0.33333),
	)
	x := NewMatrix(
		Row(9, 3, 0, 9),
		Row(-5, -2, -6, -3),
		Row(-4, 9, 6, 4),
		Row(-7, 6, 6, 2),
	)
	a := Inverse(x)
	AssertMatrixEqual(t, a, e)
}

func TestMultiplyProductByInverse(t *testing.T) {
	x := NewMatrix(
		Row(3, -9, 7, 3),
		Row(3, -8, 2, -9),
		Row(-4, 4, 4, 1),
		Row(-6, 5, -1, 1),
	)
	y := NewMatrix(
		Row(8, 2, 2, 2),
		Row(3, -1, 7, 0),
		Row(7, 0, 5, 4),
		Row(6, -2, 0, 5),
	)
	a := x.MultMat(y)
	AssertMatrixEqual(t, a.MultMat(Inverse(y)), x)
}
