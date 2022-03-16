package render

import (
	"math"
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

func TestTheHitWhenAllIntersectionsHavePositiveT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2)
	if xs.Hit() != i1 {
		t.Errorf("Hit was incorrect")
	}
}

func TestTheHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := NewIntersections(i1, i2)
	if xs.Hit() != i2 {
		t.Errorf("Hit was incorrect")
	}
}

func TestTheHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := NewIntersections(i1, i2)
	if xs.Hit() != nil {
		t.Errorf("Hit was incorrect")
	}
}

func TestTheHitIsAlwaysTheLowestNonnegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2, i3, i4)
	if xs.Hit() != i4 {
		t.Errorf("Hit was incorrect")
	}
}

func TestPrecomputingTheStateOfAnIntersection(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertFloatEqual(t, c.Distance, i.Distance)
	AssertObjectEqual(t, c.Object, i.Object)
	AssertVectorEqual(t, c.Point, alg.NewPoint(0, 0, -1))
	AssertVectorEqual(t, c.Eye, alg.NewVector3(0, 0, -1))
	AssertVectorEqual(t, c.Normal, alg.NewVector3(0, 0, -1))
}

func TestTheHitWhenAnIntersectionOccursOnTheOutside(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertFalse(t, c.Inside)
}

func TestTheHitWhenAnIntersectionOccursOnTheInside(t *testing.T) {
	r := NewRay(alg.Origin, alg.NewVector3(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(1, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertTrue(t, c.Inside)
	AssertVectorEqual(t, c.Point, alg.NewPoint(0, 0, 1))
	AssertVectorEqual(t, c.Eye, alg.NewVector3(0, 0, -1))
	AssertVectorEqual(t, c.Normal, alg.NewVector3(0, 0, -1))
}

func TestShadingAnIntersection(t *testing.T) {
	e := canvas.NewColor(0.38066, 0.47583, 0.28550)
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := w.Objects[0]
	i := NewIntersection(4, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertColorEqual(t, w.shade(c, w.MaxBounces), e)
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	e := canvas.NewColor(0.90498, 0.90498, 0.90498)
	w := newDefaultWorldWithoutLight()
	l := NewPointLight(alg.NewPoint(0, 0.25, 0), canvas.White)
	w.AddLight(l)
	r := NewRay(alg.Origin, alg.NewVector3(0, 0, 1))
	s := w.Objects[1]
	i := NewIntersection(0.5, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertColorEqual(t, w.shade(c, w.MaxBounces), e)
}

func TestTheHitShouldOffsetThePoint(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	s.Transform = alg.Translation(0, 0, 1)
	i := NewIntersection(5, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertTrue(t, c.OverPoint[2] < -EPSILON/2)
	AssertTrue(t, c.Point[2] > c.OverPoint[2])
}

func TestPrecomputingTheReflectionVector(t *testing.T) {
	f := math.Sqrt2 / 2
	r := NewRay(alg.NewPoint(0, 1, -1), alg.NewVector3(0, -f, f))
	s := NewPlane()
	i := NewIntersection(math.Sqrt2, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertVectorEqual(t, c.Reflect, alg.NewVector3(0, f, f))
}

func TestFindingN1AndN2AtVariousIntersections(t *testing.T) {
	a := newGlassSphere()
	a.Transform = alg.Scaling(2, 2, 2)
	a.Material.RefractiveIndex = 1.5
	b := newGlassSphere()
	b.Transform = alg.Translation(0, 0, -0.25)
	b.Material.RefractiveIndex = 2.0
	c := newGlassSphere()
	c.Transform = alg.Translation(0, 0, 0.25)
	c.Material.RefractiveIndex = 2.5
	r := NewRay(alg.NewPoint(0, 0, -4), alg.NewVector3(0, 0, 1))
	xs := NewIntersections(
		NewIntersection(2, a),
		NewIntersection(2.75, b),
		NewIntersection(3.25, c),
		NewIntersection(4.75, b),
		NewIntersection(5.25, c),
		NewIntersection(6, a),
	)
	ex := []struct {
		n1 float64
		n2 float64
	}{
		{n1: 1.0, n2: 1.5},
		{n1: 1.5, n2: 2.0},
		{n1: 2.0, n2: 2.5},
		{n1: 2.5, n2: 2.5},
		{n1: 2.5, n2: 1.5},
		{n1: 1.5, n2: 1.0},
	}
	for i, e := range ex {
		cp := prepareComps(xs[i], r, xs)
		AssertFloatEqual(t, cp.N1, e.n1)
		AssertFloatEqual(t, cp.N2, e.n2)
	}
}

func TestTheUnderPointIsOffsetBelowTheSurface(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := newGlassSphere()
	s.Transform = alg.Translation(0, 0, 1)
	i := NewIntersection(5, s)
	c := prepareComps(i, r, NewIntersections(i))
	AssertTrue(t, c.UnderPoint[2] > EPSILON/2)
	AssertTrue(t, c.Point[2] < c.UnderPoint[2])
}

func newGlassSphere() *Sphere {
	s := NewSphere()
	s.Material.Transparency = 1
	s.Material.RefractiveIndex = 1.5
	return s
}
