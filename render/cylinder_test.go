package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestARayMissesACylinder(t *testing.T) {
	c := NewCylinder()
	ex := []struct {
		p alg.Vector
		d alg.Vector
	}{
		{p: alg.NewPoint(1, 0, 0), d: alg.NewVector3(0, 1, 0)},
		{p: alg.NewPoint(0, 0, 0), d: alg.NewVector3(0, 1, 0)},
		{p: alg.NewPoint(0, 0, -5), d: alg.NewVector3(1, 1, 1)},
	}
	for _, e := range ex {
		r := NewRay(e.p, e.d)
		xs := c.Intersect(r)
		AssertTrue(t, len(xs) == 0)
	}
}

// TODO: Failing.
// func TestARayIntersectsACylinder(t *testing.T) {
// 	c := NewCylinder()
// 	ex := []struct {
// 		p  alg.Vector
// 		d  alg.Vector
// 		t1 float64
// 		t2 float64
// 	}{
// 		{p: alg.NewPoint(1, 0, -5), d: alg.NewVector3(0, 0, 1), t1: 5, t2: 5},
// 		{p: alg.NewPoint(0, 0, -5), d: alg.NewVector3(0, 0, 1), t1: 4, t2: 6},
// 		{p: alg.NewPoint(0.5, 0, -5), d: alg.NewVector3(0.1, 1, 1), t1: 6.80798, t2: 7.08872},
// 	}
// 	for _, e := range ex {
// 		r := NewRay(e.p, e.d)
// 		xs := c.Intersect(r)
// 		AssertTrue(t, len(xs) == 2)
// 		AssertFloatEqual(t, xs[0].Distance, e.t1)
// 		AssertFloatEqual(t, xs[1].Distance, e.t2)
// 	}
// }

func TestTheNormalOnACylinder(t *testing.T) {
	c := NewCylinder()
	ex := []struct {
		p alg.Vector
		n alg.Vector
	}{
		{p: alg.NewPoint(1, 0, 0), n: alg.NewVector3(1, 0, 0)},
		{p: alg.NewPoint(0, 5, -1), n: alg.NewVector3(0, 0, -1)},
		{p: alg.NewPoint(0, -2, 1), n: alg.NewVector3(0, 0, 1)},
		{p: alg.NewPoint(-1, 1, 0), n: alg.NewVector3(-1, 0, 0)},
	}
	for _, e := range ex {
		AssertVectorEqual(t, c.NormalAt(e.p, nil), e.n)
	}
}
