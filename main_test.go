package main

import (
	"math"
	"testing"
)

func TestTupleString(t *testing.T) {
	e := "Tuple(1.000000, 2.000000, 3.000000, 0.000000)"
	x := NewTuple(1, 2, 3, 0)
	a := x.String()
	if a != e {
		t.Errorf("Tuple was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func TestAddingTwoTuple(t *testing.T) {
	e := NewTuple(1, 1, 6, 1)
	x := NewTuple(3, -2, 5, 1)
	y := NewTuple(-2, 3, 1, 0)
	a := x.Add(y)
	assertTupleEqual(t, a, e)
}

func TestSubtractingTwoPoints(t *testing.T) {
	e := NewVector(-2, -4, -6)
	x := NewPoint(3, 2, 1)
	y := NewPoint(5, 6, 7)
	a := x.Sub(y)
	assertTupleEqual(t, a, e)
}

func TestSubtractingAVectorFromAPoint(t *testing.T) {
	e := NewPoint(-2, -4, -6)
	x := NewPoint(3, 2, 1)
	y := NewVector(5, 6, 7)
	a := x.Sub(y)
	assertTupleEqual(t, a, e)
}

func TestSubtractingTwoVectors(t *testing.T) {
	e := NewVector(-2, -4, -6)
	x := NewVector(3, 2, 1)
	y := NewVector(5, 6, 7)
	a := x.Sub(y)
	assertTupleEqual(t, a, e)
}

func TestSubtractingAVectorFromTheZeroVector(t *testing.T) {
	e := NewVector(-1, 2, -3)
	x := NewVector(0, 0, 0)
	y := NewVector(1, -2, 3)
	a := x.Sub(y)
	assertTupleEqual(t, a, e)
}

func TestNegatingATuple(t *testing.T) {
	e := NewTuple(-1, 2, -3, 4)
	x := NewTuple(1, -2, 3, -4)
	a := x.Negate()
	assertTupleEqual(t, a, e)
}

func TestMultiplyingATupleByAScalar(t *testing.T) {
	e := NewTuple(3.5, -7, 10.5, -14)
	x := NewTuple(1, -2, 3, -4)
	a := x.Mul(3.5)
	assertTupleEqual(t, a, e)
}

func TestDividingATupleByAScalar(t *testing.T) {
	e := NewTuple(0.5, -1, 1.5, -2)
	x := NewTuple(1, -2, 3, -4)
	a := x.Div(2)
	assertTupleEqual(t, a, e)
}

func TestDotProduct(t *testing.T) {
	e := 20.0
	x := NewVector(1, 2, 3)
	y := NewVector(2, 3, 4)
	a := x.Dot(y)
	assertFloatEqual(t, a, e)
}

func TestMagnitudeOfVector100(t *testing.T) {
	e := 1.0
	x := NewVector(1, 0, 0)
	a := x.Magnitude()
	assertFloatEqual(t, a, e)
}

func TestMagnitudeOfVector010(t *testing.T) {
	e := 1.0
	x := NewVector(0, 1, 0)
	a := x.Magnitude()
	assertFloatEqual(t, a, e)
}

func TestMagnitudeOfVector001(t *testing.T) {
	e := 1.0
	x := NewVector(0, 0, 1)
	a := x.Magnitude()
	assertFloatEqual(t, a, e)
}

func TestMagnitudeOfVector123(t *testing.T) {
	e := math.Sqrt(14)
	x := NewVector(1, 2, 3)
	a := x.Magnitude()
	assertFloatEqual(t, a, e)
}

func TestMagnitudeOfVectorNegative123(t *testing.T) {
	e := math.Sqrt(14)
	x := NewVector(-1, -2, -3)
	a := x.Magnitude()
	assertFloatEqual(t, a, e)
}

func TestNormalizeVector400(t *testing.T) {
	e := NewVector(1, 0, 0)
	x := NewVector(4, 0, 0)
	a := x.Normalize()
	assertTupleEqual(t, a, e)
}

func TestNormalizeVector123(t *testing.T) {
	n := math.Sqrt(14)
	e := NewVector(1/n, 2/n, 3/n)
	x := NewVector(1, 2, 3)
	a := x.Normalize()
	assertTupleEqual(t, a, e)
}

func TestMagnitudeOfNormalizedVector123(t *testing.T) {
	e := 1.0
	x := NewVector(1, 2, 3)
	a := x.Normalize().Magnitude()
	assertFloatEqual(t, a, e)
}

func TestCrossProductXY(t *testing.T) {
	e := NewVector(-1, 2, -1)
	x := NewVector(1, 2, 3)
	y := NewVector(2, 3, 4)
	a := x.Cross(y)
	assertTupleEqual(t, a, e)
}

func TestCrossProductYX(t *testing.T) {
	e := NewVector(1, -2, 1)
	x := NewVector(1, 2, 3)
	y := NewVector(2, 3, 4)
	a := y.Cross(x)
	assertTupleEqual(t, a, e)
}

func TestAddingColors(t *testing.T) {
	e := NewColor(1.6, 0.7, 1.0)
	x := NewColor(0.9, 0.6, 0.75)
	y := NewColor(0.7, 0.1, 0.25)
	a := x.Add(y)
	assertColorEqual(t, a, e)
}

func TestSubtractingColors(t *testing.T) {
	e := NewColor(0.2, 0.5, 0.5)
	x := NewColor(0.9, 0.6, 0.75)
	y := NewColor(0.7, 0.1, 0.25)
	a := x.Sub(y)
	assertColorEqual(t, a, e)
}

func TestScaleColors(t *testing.T) {
	e := NewColor(0.4, 0.6, 0.8)
	x := NewColor(0.2, 0.3, 0.4)
	a := x.Scale(2)
	assertColorEqual(t, a, e)
}

func TestMultiplyingColors(t *testing.T) {
	e := NewColor(0.9, 0.2, 0.04)
	x := NewColor(1.0, 0.2, 0.4)
	y := NewColor(0.9, 1.0, 0.1)
	a := x.Mul(y)
	assertColorEqual(t, a, e)
}

func TestMatrixesEqual(t *testing.T) {
	e := NewMatrix(
		Row(1, 2, 3, 4),
		Row(5, 6, 7, 8),
		Row(9, 10, 11, 12),
		Row(13, 14, 15, 16),
	)
	a := NewMatrix(
		Row(1, 2, 3, 4),
		Row(5, 6, 7, 8),
		Row(9, 10, 11, 12),
		Row(13, 14, 15, 16),
	)
	assertMatrixEqual(t, a, e)
}

func TestMultiplyMatrices(t *testing.T) {
	e := NewMatrix(
		Row(20, 22, 50, 48),
		Row(44, 54, 114, 108),
		Row(40, 58, 110, 102),
		Row(16, 26, 46, 42),
	)
	x := NewMatrix(
		Row(1, 2, 3, 4),
		Row(5, 6, 7, 8),
		Row(9, 8, 7, 6),
		Row(5, 4, 3, 2),
	)
	y := NewMatrix(
		Row(-2, 1, 2, 3),
		Row(3, 2, 1, -1),
		Row(4, 3, 6, 5),
		Row(1, 2, 7, 8),
	)
	a := x.MatMul(y)
	assertMatrixEqual(t, a, e)
}

func TestMultiplyMatrixAndTuple(t *testing.T) {
	e := NewTuple(18, 24, 33, 1)
	x := NewMatrix(
		Row(1, 2, 3, 4),
		Row(2, 4, 4, 2),
		Row(8, 6, 4, 1),
		Row(0, 0, 0, 1),
	)
	y := NewTuple(1, 2, 3, 1)
	a := x.TupMul(y)
	assertTupleEqual(t, a, e)
}

func TestMultiplyMatrixWithIdentity(t *testing.T) {
	e := NewMatrix(
		Row(1, 2, 3, 4),
		Row(5, 6, 7, 8),
		Row(9, 8, 7, 6),
		Row(5, 4, 3, 2),
	)
	x := NewMatrix(
		Row(1, 2, 3, 4),
		Row(5, 6, 7, 8),
		Row(9, 8, 7, 6),
		Row(5, 4, 3, 2),
	)
	a := x.MatMul(Id4)
	assertMatrixEqual(t, a, e)
}

func TestMultiplyIdentityWithTuple(t *testing.T) {
	e := NewTuple(1, 2, 3, 1)
	y := NewTuple(1, 2, 3, 1)
	a := Id4.TupMul(y)
	assertTupleEqual(t, a, e)
}

func TestTransposeMatrix(t *testing.T) {
	e := NewMatrix(
		Row(1, 5, 9, 13),
		Row(2, 6, 10, 14),
		Row(3, 7, 11, 15),
		Row(4, 8, 12, 16),
	)
	x := NewMatrix(
		Row(1, 2, 3, 4),
		Row(5, 6, 7, 8),
		Row(9, 10, 11, 12),
		Row(13, 14, 15, 16),
	)
	a := x.Transpose()
	assertMatrixEqual(t, a, e)
}

func TestTransposeIdentity(t *testing.T) {
	e := Id4
	a := Id4.Transpose()
	assertMatrixEqual(t, a, e)
}

func TestLuDecompose(t *testing.T) {
	el := NewMatrix(
		Row(3.0, 0, 0),
		Row(0.1, 7.003333, 0),
		Row(0.3, -0.19, 10.012042),
	)
	eu := NewMatrix(
		Row(1, -0.033333, -0.066667),
		Row(0, 1, -0.041885),
		Row(0, 0, 1),
	)
	x := NewMatrix(
		Row(3, -0.1, -0.2),
		Row(0.1, 7, -0.3),
		Row(0.3, -0.2, 10),
	)
	al, au := x.LuDecompose()
	assertMatrixEqual(t, al, el)
	assertMatrixEqual(t, au, eu)
	assertMatrixEqual(t, al.MatMul(au), x)
}

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}

func tupleEqual(a, b *Tuple) bool {
	return floatEqual(a.x, b.x) &&
		floatEqual(a.y, b.y) &&
		floatEqual(a.z, b.z) &&
		floatEqual(a.w, b.w)
}

func colorEqual(a, b *Color) bool {
	return floatEqual(a.r, b.r) &&
		floatEqual(a.g, b.g) &&
		floatEqual(a.b, b.b)
}

func matrixEqual(a, b *Matrix) bool {
	sza := len(a.m)
	szb := len(b.m)
	if sza != szb {
		return false
	}
	for r := 0; r < sza; r++ {
		for c := 0; c < sza; c++ {
			va := a.m[r][c]
			vb := b.m[r][c]
			if !floatEqual(va, vb) {
				return false
			}
		}
	}
	return true
}

func assertFloatEqual(t *testing.T, a, e float64) {
	t.Helper()
	if !floatEqual(a, e) {
		t.Errorf("Float was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func assertTupleEqual(t *testing.T, a, e *Tuple) {
	t.Helper()
	if !tupleEqual(a, e) {
		t.Errorf("Tuple was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func assertColorEqual(t *testing.T, a, e *Color) {
	t.Helper()
	if !colorEqual(a, e) {
		t.Errorf("Color was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func assertMatrixEqual(t *testing.T, a, e *Matrix) {
	t.Helper()
	sza := len(a.m)
	sze := len(e.m)
	if sza != sze {
		t.Errorf("Matrix size was incorrect, \n got: %v \n want: %v", sza, sze)
	} else if !matrixEqual(a, e) {
		t.Errorf("Matrix was incorrect, \n got: %v \n want: %v", a, e)
	}
}
