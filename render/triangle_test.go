package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestConstructingATriangle(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	AssertVectorEqual(t, tr.P1, p1)
	AssertVectorEqual(t, tr.P2, p2)
	AssertVectorEqual(t, tr.P3, p3)
	AssertVectorEqual(t, tr.E1, alg.NewVector3(-1, -1, 0))
	AssertVectorEqual(t, tr.E2, alg.NewVector3(1, -1, 0))
	AssertVectorEqual(t, tr.Normal, alg.NewVector3(0, 0, -1))
}

func TestFindingTheNormalOnATriangle(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	AssertVectorEqual(t, tr.Normal, tr.NormalAt(alg.NewPoint(0, 0.5, 0), nil))
	AssertVectorEqual(t, tr.Normal, tr.NormalAt(alg.NewPoint(-0.5, 0.75, 0), nil))
	AssertVectorEqual(t, tr.Normal, tr.NormalAt(alg.NewPoint(0.5, 0.25, 0), nil))
}

func TestIntersectRayParallelToTriangle(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	r := NewRay(alg.NewPoint(0, -1, -2), alg.NewVector3(0, 1, 0))
	xs := tr.Intersect(r)
	AssertTrue(t, len(xs) == 0)
}

func TestARayMissesP1P3Edge(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	r := NewRay(alg.NewPoint(1, 1, -2), alg.NewVector3(0, 0, 1))
	xs := tr.Intersect(r)
	AssertTrue(t, len(xs) == 0)
}

func TestARayMissesP1P2Edge(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	r := NewRay(alg.NewPoint(-1, 1, -2), alg.NewVector3(0, 0, 1))
	xs := tr.Intersect(r)
	AssertTrue(t, len(xs) == 0)
}

func TestARayMissesP2P3Edge(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	r := NewRay(alg.NewPoint(0, -1, -2), alg.NewVector3(0, 0, 1))
	xs := tr.Intersect(r)
	AssertTrue(t, len(xs) == 0)
}

func TestARayStrikesATriangle(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	r := NewRay(alg.NewPoint(0, 0.5, -2), alg.NewVector3(0, 0, 1))
	xs := tr.Intersect(r)
	AssertTrue(t, len(xs) == 1)
	AssertFloatEqual(t, xs[0].Distance, 2)
}
