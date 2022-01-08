package alg

import (
	"testing"
)

func TestDeterminantOf2x2Matrix(t *testing.T) {
	a := NewMatrix(
		Row(1, 5),
		Row(-3, 2),
	)
	AssertFloatEqual(t, Det(a), 17)
}

func TestSubMatrix2x2(t *testing.T) {
	es := []struct {
		r int
		c int
		m Matrix
	}{
		{0, 0, NewMatrix(Row(2, 7), Row(6, -3))},
		{0, 1, NewMatrix(Row(-3, 7), Row(0, -3))},
		{0, 2, NewMatrix(Row(-3, 2), Row(0, 6))},
		{1, 0, NewMatrix(Row(5, 0), Row(6, -3))},
		{1, 1, NewMatrix(Row(1, 0), Row(0, -3))},
		{1, 2, NewMatrix(Row(1, 5), Row(0, 6))},
		{2, 0, NewMatrix(Row(5, 0), Row(2, 7))},
		{2, 1, NewMatrix(Row(1, 0), Row(-3, 7))},
		{2, 2, NewMatrix(Row(1, 5), Row(-3, 2))},
	}
	a := NewMatrix(
		Row(1, 5, 0),
		Row(-3, 2, 7),
		Row(0, 6, -3),
	)
	for _, e := range es {
		AssertMatrixEqual(t, SubMatrix(a, e.r, e.c), e.m)
	}
}

func TestSubMatrix3x3(t *testing.T) {
	es := []struct {
		r int
		c int
		m Matrix
	}{
		{2, 1, NewMatrix(Row(-6, 1, 6), Row(-8, 8, 6), Row(-7, -1, 1))},
	}
	a := NewMatrix(
		Row(-6, 1, 1, 6),
		Row(-8, 5, 8, 6),
		Row(-1, 0, 8, 2),
		Row(-7, 1, -1, 1),
	)
	for _, e := range es {
		AssertMatrixEqual(t, SubMatrix(a, e.r, e.c), e.m)
	}
}

func TestMinor(t *testing.T) {
	m := NewMatrix(
		Row(3, 5, 0),
		Row(2, -1, -7),
		Row(6, -1, 5),
	)
	AssertFloatEqual(t, Det(SubMatrix(m, 1, 0)), 25)
	AssertFloatEqual(t, Minor(m, 1, 0), 25)
}

func TestMinor2(t *testing.T) {
	m := NewMatrix(
		Row(3, 5, 0),
		Row(2, -1, -7),
		Row(6, -1, 5),
	)
	AssertFloatEqual(t, Minor(m, 0, 0), -12)
	AssertFloatEqual(t, Cofactor(m, 0, 0), -12)
	AssertFloatEqual(t, Minor(m, 1, 0), 25)
	AssertFloatEqual(t, Cofactor(m, 1, 0), -25)
}

func TestDeterminantOf3x3Matrix(t *testing.T) {
	a := NewMatrix(
		Row(1, 2, 6),
		Row(-5, 8, -4),
		Row(2, 6, 4),
	)
	AssertFloatEqual(t, Cofactor(a, 0, 0), 56)
	AssertFloatEqual(t, Cofactor(a, 0, 1), 12)
	AssertFloatEqual(t, Cofactor(a, 0, 2), -46)
	AssertFloatEqual(t, Det(a), -196)
}

func TestDeterminantOf4x4Matrix(t *testing.T) {
	a := NewMatrix(
		Row(-2, -8, 3, 5),
		Row(-3, 1, 7, 3),
		Row(1, 2, -9, 6),
		Row(-6, 7, 7, -9),
	)
	AssertFloatEqual(t, Cofactor(a, 0, 0), 690)
	AssertFloatEqual(t, Cofactor(a, 0, 1), 447)
	AssertFloatEqual(t, Cofactor(a, 0, 2), 210)
	AssertFloatEqual(t, Cofactor(a, 0, 3), 51)
	AssertFloatEqual(t, Det(a), -4071)
}

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
