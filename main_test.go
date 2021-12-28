package main

import (
	"math"
	"testing"
)

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

const float64EqualityThreshold = 1e-6

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
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
		t.Errorf("Sum was incorrect, \n got: %v \n want: %v", a, e)
	}
}
