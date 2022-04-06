package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestIntersectingARayWithAnEmptyGroup(t *testing.T) {
	g := NewGroup()
	r := NewRay(alg.NewPoint(0, 0, 0), alg.NewVector3(0, 0, 1))
	xs := g.Intersect(r)
	AssertIntEqual(t, len(xs), 0)
}

func TestIntersectingARayWithANonEmptyGroup(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()
	s2 := NewSphere()
	s2.Transform = alg.Translation(0, 0, -3)
	s3 := NewSphere()
	s3.Transform = alg.Translation(5, 0, 0)
	g.AddKid(s1)
	g.AddKid(s2)
	g.AddKid(s3)
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	xs := g.Intersect(r)
	AssertIntEqual(t, len(xs), 4)
	AssertObjectEqual(t, xs[0].Object, s2)
	AssertObjectEqual(t, xs[1].Object, s2)
	AssertObjectEqual(t, xs[2].Object, s1)
	AssertObjectEqual(t, xs[3].Object, s1)
}

func TestIntersectingATransformedGroup(t *testing.T) {
	g := NewGroup()
	g.Transform = alg.Scaling(2, 2, 2)
	s := NewSphere()
	s.Transform = alg.Translation(5, 0, 0)
	g.AddKid(s)
	r := NewRay(alg.NewPoint(10, 0, -10), alg.NewVector3(0, 0, 1))
	xs := g.Intersect(r)
	AssertIntEqual(t, len(xs), 2)
}
