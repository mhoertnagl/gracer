package main

import (
	"math"
	"testing"
)

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

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}

func colorEqual(a, b *Color) bool {
	return floatEqual(a.r, b.r) &&
		floatEqual(a.g, b.g) &&
		floatEqual(a.b, b.b)
}

func assertColorEqual(t *testing.T, a, e *Color) {
	t.Helper()
	if !colorEqual(a, e) {
		t.Errorf("Color was incorrect, \n got: %v \n want: %v", a, e)
	}
}
