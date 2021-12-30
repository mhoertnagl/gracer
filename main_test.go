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

// func TestAddingTwoTuple(t *testing.T) {
// 	e := tuple(1, 1, 6, 1)
// 	x := tuple(3, -2, 5, 1)
// 	y := tuple(-2, 3, 1, 0)
// 	a := add(x, y)
// 	assertTupleEqual(t, a, e)
// }

// func TestSubtractingTwoPoints(t *testing.T) {
// 	e := vector(-2, -4, -6)
// 	x := point(3, 2, 1)
// 	y := point(5, 6, 7)
// 	a := sub(x, y)
// 	assertTupleEqual(t, a, e)
// }

// func TestSubtractingAVectorFromAPoint(t *testing.T) {
// 	e := point(-2, -4, -6)
// 	x := point(3, 2, 1)
// 	y := vector(5, 6, 7)
// 	a := sub(x, y)
// 	assertTupleEqual(t, a, e)
// }

// func TestSubtractingTwoVectors(t *testing.T) {
// 	e := vector(-2, -4, -6)
// 	x := vector(3, 2, 1)
// 	y := vector(5, 6, 7)
// 	a := sub(x, y)
// 	assertTupleEqual(t, a, e)
// }

// func TestSubtractingAVectorFromTheZeroVector(t *testing.T) {
// 	e := vector(-1, 2, -3)
// 	x := vector(0, 0, 0)
// 	y := vector(1, -2, 3)
// 	a := sub(x, y)
// 	assertTupleEqual(t, a, e)
// }

// func TestNegatingATuple(t *testing.T) {
// 	e := tuple(-1, 2, -3, 4)
// 	x := tuple(1, -2, 3, -4)
// 	a := negate(x)
// 	assertTupleEqual(t, a, e)
// }

// func TestMultiplyingATupleByAScalar(t *testing.T) {
// 	e := tuple(3.5, -7, 10.5, -14)
// 	x := tuple(1, -2, 3, -4)
// 	a := mul(3.5, x)
// 	assertTupleEqual(t, a, e)
// }

// func TestDividingATupleByAScalar(t *testing.T) {
// 	e := tuple(0.5, -1, 1.5, -2)
// 	x := tuple(1, -2, 3, -4)
// 	a := div(x, 2)
// 	assertTupleEqual(t, a, e)
// }

// func TestDotProduct(t *testing.T) {
// 	e := 20.0
// 	x := vector(1, 2, 3)
// 	y := vector(2, 3, 4)
// 	a := dot(x, y)
// 	assertFloatEqual(t, a, e)
// }

// func TestMagnitudeOfVector100(t *testing.T) {
// 	e := 1.0
// 	x := vector(1, 0, 0)
// 	a := magnitude(x)
// 	assertFloatEqual(t, a, e)
// }

// func TestMagnitudeOfVector010(t *testing.T) {
// 	e := 1.0
// 	x := vector(0, 1, 0)
// 	a := magnitude(x)
// 	assertFloatEqual(t, a, e)
// }

// func TestMagnitudeOfVector001(t *testing.T) {
// 	e := 1.0
// 	x := vector(0, 0, 1)
// 	a := magnitude(x)
// 	assertFloatEqual(t, a, e)
// }

// func TestMagnitudeOfVector123(t *testing.T) {
// 	e := math.Sqrt(14)
// 	x := vector(1, 2, 3)
// 	a := magnitude(x)
// 	assertFloatEqual(t, a, e)
// }

// func TestMagnitudeOfVectorNegative123(t *testing.T) {
// 	e := math.Sqrt(14)
// 	x := vector(-1, -2, -3)
// 	a := magnitude(x)
// 	assertFloatEqual(t, a, e)
// }

// func TestNormalizeVector400(t *testing.T) {
// 	e := vector(1, 0, 0)
// 	x := vector(4, 0, 0)
// 	a := normalize(x)
// 	assertTupleEqual(t, a, e)
// }

// func TestNormalizeVector123(t *testing.T) {
// 	n := math.Sqrt(14)
// 	e := vector(1/n, 2/n, 3/n)
// 	x := vector(1, 2, 3)
// 	a := normalize(x)
// 	assertTupleEqual(t, a, e)
// }

// func TestMagnitudeOfNormalizedVector123(t *testing.T) {
// 	e := 1.0
// 	x := vector(1, 2, 3)
// 	a := magnitude(normalize(x))
// 	assertFloatEqual(t, a, e)
// }

// func TestCrossProductXY(t *testing.T) {
// 	e := vector(-1, 2, -1)
// 	x := vector(1, 2, 3)
// 	y := vector(2, 3, 4)
// 	a := cross(x, y)
// 	assertTupleEqual(t, a, e)
// }

// func TestCrossProductYX(t *testing.T) {
// 	e := vector(1, -2, 1)
// 	x := vector(1, 2, 3)
// 	y := vector(2, 3, 4)
// 	a := cross(y, x)
// 	assertTupleEqual(t, a, e)
// }

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

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}

func tupleEqual(a, b *Tuple) bool {
	return floatEqual(a.x, b.x) &&
		floatEqual(a.y, b.y) &&
		floatEqual(a.z, b.z) &&
		floatEqual(a.w, b.w)
}

func assertTupleEqual(t *testing.T, a, e *Tuple) {
	t.Helper()
	if !tupleEqual(a, e) {
		t.Errorf("Tuple was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func assertFloatEqual(t *testing.T, a, e float64) {
	t.Helper()
	if !floatEqual(a, e) {
		t.Errorf("Float was incorrect, \n got: %v \n want: %v", a, e)
	}
}
