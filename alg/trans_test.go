package alg

import (
	"math"
	"testing"
)

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

func TestRotatingHalfQuarterAroundX(t *testing.T) {
	e := NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)
	tr := RotationX(math.Pi / 4)
	p := NewPoint(0, 1, 0)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingFullQuarterAroundX(t *testing.T) {
	e := NewPoint(0, 0, 1)
	tr := RotationX(math.Pi / 2)
	p := NewPoint(0, 1, 0)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingInverseHalfQuarterAroundX(t *testing.T) {
	e := NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)
	tr := RotationX(math.Pi / 4)
	inv := Inverse(tr)
	p := NewPoint(0, 1, 0)
	a := inv.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingHalfQuarterAroundY(t *testing.T) {
	e := NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)
	tr := RotationY(math.Pi / 4)
	p := NewPoint(0, 0, 1)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingFullQuarterAroundY(t *testing.T) {
	e := NewPoint(1, 0, 0)
	tr := RotationY(math.Pi / 2)
	p := NewPoint(0, 0, 1)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingInverseHalfQuarterAroundY(t *testing.T) {
	e := NewPoint(-math.Sqrt2/2, 0, math.Sqrt2/2)
	tr := RotationY(math.Pi / 4)
	inv := Inverse(tr)
	p := NewPoint(0, 0, 1)
	a := inv.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingHalfQuarterAroundZ(t *testing.T) {
	e := NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0)
	tr := RotationZ(math.Pi / 4)
	p := NewPoint(0, 1, 0)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingFullQuarterAroundZ(t *testing.T) {
	e := NewPoint(-1, 0, 0)
	tr := RotationZ(math.Pi / 2)
	p := NewPoint(0, 1, 0)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestRotatingInverseHalfQuarterAroundZ(t *testing.T) {
	e := NewPoint(math.Sqrt2/2, math.Sqrt2/2, 0)
	tr := RotationZ(math.Pi / 4)
	inv := Inverse(tr)
	p := NewPoint(0, 1, 0)
	a := inv.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestShearingXInProportionToY(t *testing.T) {
	e := NewPoint(5, 3, 4)
	tr := Shearing(1, 0, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestShearingXInProportionToZ(t *testing.T) {
	e := NewPoint(6, 3, 4)
	tr := Shearing(0, 1, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestShearingYInProportionToX(t *testing.T) {
	e := NewPoint(2, 5, 4)
	tr := Shearing(0, 0, 1, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestShearingYInProportionToZ(t *testing.T) {
	e := NewPoint(2, 7, 4)
	tr := Shearing(0, 0, 0, 1, 0, 0)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestShearingZInProportionToX(t *testing.T) {
	e := NewPoint(2, 3, 6)
	tr := Shearing(0, 0, 0, 0, 1, 0)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestShearingZInProportionToY(t *testing.T) {
	e := NewPoint(2, 3, 7)
	tr := Shearing(0, 0, 0, 0, 0, 1)
	p := NewPoint(2, 3, 4)
	a := tr.MultVec(p)
	AssertVectorEqual(t, a, e)
}

func TestTransformationSequenceApplication(t *testing.T) {
	p := NewPoint(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)
	p2 := A.MultVec(p)
	AssertVectorEqual(t, p2, NewPoint(1, -1, 0))
	p3 := B.MultVec(p2)
	AssertVectorEqual(t, p3, NewPoint(5, -5, 0))
	p4 := C.MultVec(p3)
	AssertVectorEqual(t, p4, NewPoint(15, 0, 7))
}

func TestTransformationManualChaining(t *testing.T) {
	p := NewPoint(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)
	T := C.MultMat(B).MultMat(A)
	AssertVectorEqual(t, T.MultVec(p), NewPoint(15, 0, 7))
}

func TestTransformationDotChaining(t *testing.T) {
	p := NewPoint(1, 0, 1)
	T := Id4.
		RotateX(math.Pi/2).
		Scale(5, 5, 5).
		Translate(10, 5, 7)
	AssertVectorEqual(t, T.MultVec(p), NewPoint(15, 0, 7))
}
