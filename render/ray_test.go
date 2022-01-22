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
	AssertFloatEqual(t, xs[0].t, 4.0)
	AssertFloatEqual(t, xs[1].t, 6.0)
}

func TestARayIntersectsASphereAtATangent(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 1, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].t, 5.0)
	AssertFloatEqual(t, xs[1].t, 5.0)
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
	AssertFloatEqual(t, xs[0].t, -1.0)
	AssertFloatEqual(t, xs[1].t, 1.0)
}

func TestARayBehindASphere(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, 5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	xs := s.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
	AssertFloatEqual(t, xs[0].t, -6.0)
	AssertFloatEqual(t, xs[1].t, -4.0)
}
