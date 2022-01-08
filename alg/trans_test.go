package alg

import "testing"

func TestMultiplyByTranslation(t *testing.T) {
	e := NewPoint(2, 1, 7)
	tr := Translation(5, -3, 2)
	p := NewPoint(-3, 4, 5)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestMultiplyByInverseTranslation(t *testing.T) {
	e := NewPoint(-8, 7, 3)
	tr := Translation(5, -3, 2)
	inv := Inverse(tr)
	p := NewPoint(-3, 4, 5)
	a := inv.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	tr := Translation(5, -3, 2)
	v := NewVector3(-3, 4, 5)
	AssertVectorEqual(t, tr.MultVec(v), v)
}

func TestMultiplyPointByScaling(t *testing.T) {
	e := NewPoint(-8, 18, 32)
	tr := Scaling(2, 3, 4)
	p := NewPoint(-4, 6, 8)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestMultiplyVectorByScaling(t *testing.T) {
	e := NewVector3(-8, 18, 32)
	tr := Scaling(2, 3, 4)
	p := NewVector3(-4, 6, 8)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestMultiplyByInverseScaling(t *testing.T) {
	e := NewVector3(-2, 2, 2)
	tr := Scaling(2, 3, 4)
	inv := Inverse(tr)
	v := NewVector3(-4, 6, 8)
	a := inv.MultVec(v)
	AssertVectorEqual(t, a, e)
}

func TestReflectPointByScaling(t *testing.T) {
	e := NewPoint(-2, 3, 4)
	tr := Scaling(-1, 1, 1)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}
