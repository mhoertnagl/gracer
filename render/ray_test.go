package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestComputingAPointFromADistance(t *testing.T) {
	r := NewRay(alg.NewPoint(2, 3, 4), alg.NewVector3(1, 0, 0))
	AssertVectorEqual(t, r.Position(0), alg.NewPoint(2, 3, 4))
	AssertVectorEqual(t, r.Position(1), alg.NewPoint(3, 3, 4))
	AssertVectorEqual(t, r.Position(-1), alg.NewPoint(1, 3, 4))
	AssertVectorEqual(t, r.Position(2.5), alg.NewPoint(4.5, 3, 4))
}

func TestARayIntersectsASphereAtTwoPoints(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].Distance, 4.0)
	AssertFloatEqual(t, xs[1].Distance, 6.0)
}

func TestARayIntersectsASphereAtATangent(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 1, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].Distance, 5.0)
	AssertFloatEqual(t, xs[1].Distance, 5.0)
}

func TestARayMissesASphere(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 2, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 0)
}

func TestARayOriginatesInsideASphere(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, 0), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].Distance, -1.0)
	AssertFloatEqual(t, xs[1].Distance, 1.0)
}

func TestARayBehindASphere(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, 5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].Distance, -6.0)
	AssertFloatEqual(t, xs[1].Distance, -4.0)
}

func TestTranslatingARay(t *testing.T) {
	r := NewRay(alg.NewPoint(1, 2, 3), alg.NewVector3(0, 1, 0))
	m := alg.Translation(3, 4, 5)
	r2 := r.Transform(m)
	AssertVectorEqual(t, r2.Origin, alg.NewPoint(4, 6, 8))
	AssertVectorEqual(t, r2.Direction, alg.NewVector3(0, 1, 0))
}

func TestScalingARay(t *testing.T) {
	r := NewRay(alg.NewPoint(1, 2, 3), alg.NewVector3(0, 1, 0))
	m := alg.Scaling(2, 3, 4)
	r2 := r.Transform(m)
	AssertVectorEqual(t, r2.Origin, alg.NewPoint(2, 6, 12))
	AssertVectorEqual(t, r2.Direction, alg.NewVector3(0, 3, 0))
}

func TestIntersectingAScaledSphereWithARay(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	s.Transform = alg.Scaling(2, 2, 2)
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].Distance, 3)
	AssertFloatEqual(t, xs[1].Distance, 7)
}

func TestIntersectingATranslatedSphereWithARay(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	s.Transform = alg.Translation(5, 0, 0)
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 0)
}
