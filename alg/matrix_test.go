package alg

import "testing"

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
	AssertMatrixEqual(t, a, e)
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
	a := x.MultMat(y)
	AssertMatrixEqual(t, a, e)
}

func TestMultiplyMatrixAndTuple(t *testing.T) {
	e := NewVector(18, 24, 33, 1)
	x := NewMatrix(
		Row(1, 2, 3, 4),
		Row(2, 4, 4, 2),
		Row(8, 6, 4, 1),
		Row(0, 0, 0, 1),
	)
	y := NewVector(1, 2, 3, 1)
	a := x.MultVec(y)
	AssertVectorEqual(t, a, e)
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
	a := x.MultMat(Id4)
	AssertMatrixEqual(t, a, e)
}

func TestMultiplyIdentityWithTuple(t *testing.T) {
	e := NewVector(1, 2, 3, 1)
	y := NewVector(1, 2, 3, 1)
	a := Id4.MultVec(y)
	AssertVectorEqual(t, a, e)
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
	AssertMatrixEqual(t, a, e)
}

func TestTransposeIdentity(t *testing.T) {
	e := Id4
	a := Id4.Transpose()
	AssertMatrixEqual(t, a, e)
}
