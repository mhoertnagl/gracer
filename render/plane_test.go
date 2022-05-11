package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestTheNormalVectorOfAPlane(t *testing.T) {
	p := NewPlane()
	AssertVectorEqual(t, p.NormalAt(alg.Origin, nil), alg.NewVector3(0, 1, 0))
	AssertVectorEqual(t, p.NormalAt(alg.NewPoint(10, 0, -10), nil), alg.NewVector3(0, 1, 0))
	AssertVectorEqual(t, p.NormalAt(alg.NewPoint(-5, 0, 150), nil), alg.NewVector3(0, 1, 0))
}

func TestIntersectRayParallelToPlane(t *testing.T) {
	p := NewPlane()
	r := NewRay(alg.NewPoint(0, 10, 0), alg.NewVector3(0, 0, 1))
	xs := p.Intersect(r)
	AssertIntEqual(t, len(xs), 0)
}

func TestIntersectPlaneWithCoplanarRay(t *testing.T) {
	p := NewPlane()
	r := NewRay(alg.Origin, alg.NewVector3(0, 0, 1))
	xs := p.Intersect(r)
	AssertIntEqual(t, len(xs), 0)
}

func TestARayIntersectingAPlaneFromAbove(t *testing.T) {
	p := NewPlane()
	r := NewRay(alg.NewPoint(0, 1, 0), alg.NewVector3(0, -1, 0))
	xs := p.Intersect(r)
	AssertIntEqual(t, len(xs), 1)
	AssertIntEqual(t, int(xs[0].Distance), 1)
	AssertObjectEqual(t, xs[0].Object, p)
}

func TestARayIntersectingAPlaneFromBelow(t *testing.T) {
	p := NewPlane()
	r := NewRay(alg.NewPoint(0, -1, 0), alg.NewVector3(0, 1, 0))
	xs := p.Intersect(r)
	AssertIntEqual(t, len(xs), 1)
	AssertIntEqual(t, int(xs[0].Distance), 1)
	AssertObjectEqual(t, xs[0].Object, p)
}

func TestARayIntersectingARotatedPlane(t *testing.T) {
	p := NewPlane()
	p.Transform = alg.Translation(0, -1, 0)
	r := NewRay(alg.NewPoint(0, 1, 0), alg.NewVector3(0, -1, 0))
	xs := p.Intersect(r)
	AssertIntEqual(t, len(xs), 1)
	AssertIntEqual(t, int(xs[0].Distance), 2)
	AssertObjectEqual(t, xs[0].Object, p)
}
