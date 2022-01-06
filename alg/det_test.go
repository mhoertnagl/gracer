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
