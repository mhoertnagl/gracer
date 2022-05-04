package alg

import (
	"math"
	"testing"
)

func TestVectorString(t *testing.T) {
	e := "Vector(1.000000, 2.000000, 3.000000, 0.000000)"
	x := NewVector(1, 2, 3, 0)
	a := x.String()
	if a != e {
		t.Errorf("Vector was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func TestVectorAddition(t *testing.T) {
	e := NewVector(1, 1, 6, 1)
	x := NewVector(3, -2, 5, 1)
	y := NewVector(-2, 3, 1, 0)
	a := x.Add(y)
	AssertVectorEqual(t, a, e)
}

func TestVectorSubtraction(t *testing.T) {
	e := NewVector(-2, -4, -6, 0)
	x := NewVector(3, 2, 1, 1)
	y := NewVector(5, 6, 7, 1)
	a := x.Sub(y)
	AssertVectorEqual(t, a, e)
}

func TestVectorNegation(t *testing.T) {
	e := NewVector(-1, 2, -3, 4)
	x := NewVector(1, -2, 3, -4)
	a := x.Neg()
	AssertVectorEqual(t, a, e)
}

func TestVectorScalarMultiplication(t *testing.T) {
	e := NewVector(3.5, -7, 10.5, -14)
	x := NewVector(1, -2, 3, -4)
	a := x.Mult(3.5)
	AssertVectorEqual(t, a, e)
}

func TestVectorScalarDivision(t *testing.T) {
	e := NewVector(0.5, -1, 1.5, -2)
	x := NewVector(1, -2, 3, -4)
	a := x.Div(2)
	AssertVectorEqual(t, a, e)
}

func TestVectorDotProduct(t *testing.T) {
	e := 20.0
	x := NewVector(1, 2, 3)
	y := NewVector(2, 3, 4)
	a := x.Dot(y)
	AssertFloatEqual(t, a, e)
}

func TestCrossProduct(t *testing.T) {
	e := NewVector3(1, -2, 1)
	x := NewVector3(1, 2, 3)
	y := NewVector3(2, 3, 4)
	a := y.Cross(x)
	AssertVectorEqual(t, a, e)
}

func TestVectorHadamardProduct(t *testing.T) {
	e := NewVector(2, -6, 12)
	x := NewVector(1, 2, -3)
	y := NewVector(2, -3, -4)
	a := x.Had(y)
	AssertVectorEqual(t, a, e)
}

func TestVectorMagnitude(t *testing.T) {
	ds := []struct {
		e float64
		v Vector
	}{
		{1.0, NewVector(1, 0, 0)},
		{1.0, NewVector(0, 1, 0)},
		{1.0, NewVector(0, 0, 1)},
		{math.Sqrt(14), NewVector(1, 2, 3)},
		{math.Sqrt(14), NewVector(-1, -2, -3)},
	}
	for _, d := range ds {
		AssertFloatEqual(t, d.v.Mag(), d.e)
	}
}

func TestVectorNormalization(t *testing.T) {
	n := math.Sqrt(14)
	ds := []struct {
		e Vector
		v Vector
	}{
		{NewVector(1, 0, 0), NewVector(4, 0, 0)},
		{NewVector(1/n, 2/n, 3/n), NewVector(1, 2, 3)},
	}
	for _, d := range ds {
		AssertVectorEqual(t, d.v.Norm(), d.e)
	}
}

func TestNormalizedVectorMagnitude(t *testing.T) {
	ds := []struct {
		e float64
		v Vector
	}{
		{1.0, NewVector(1, 2, 3)},
	}
	for _, d := range ds {
		AssertFloatEqual(t, d.v.Norm().Mag(), d.e)
	}
}

func TestReflectVectorAt45Deg(t *testing.T) {
	v := NewVector(1, -1, 0)
	n := NewVector(0, 1, 0)
	AssertVectorEqual(t, v.Reflect(n), NewVector(1, 1, 0))
}

func TestReflectVectorOffASlantedSurface(t *testing.T) {
	f := math.Sqrt(2) / 2
	v := NewVector(0, -1, 0)
	n := NewVector(f, f, 0)
	AssertVectorEqual(t, v.Reflect(n), NewVector(1, 0, 0))
}
