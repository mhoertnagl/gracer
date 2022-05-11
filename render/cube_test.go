package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestARayIntersectsACube(t *testing.T) {
	c := NewCube()
	ex := []struct {
		p  alg.Vector
		d  alg.Vector
		t1 float64
		t2 float64
	}{
		{p: alg.NewPoint(5, 0.5, 0), d: alg.NewVector3(-1, 0, 0), t1: 4, t2: 6},
		{p: alg.NewPoint(-5, 0.5, 0), d: alg.NewVector3(1, 0, 0), t1: 4, t2: 6},
		{p: alg.NewPoint(0.5, 5, 0), d: alg.NewVector3(0, -1, 0), t1: 4, t2: 6},
		{p: alg.NewPoint(0.5, -5, 0), d: alg.NewVector3(0, 1, 0), t1: 4, t2: 6},
		{p: alg.NewPoint(0.5, 0, 5), d: alg.NewVector3(0, 0, -1), t1: 4, t2: 6},
		{p: alg.NewPoint(0.5, 0, -5), d: alg.NewVector3(0, 0, 1), t1: 4, t2: 6},
		{p: alg.NewPoint(0, 0.5, 0), d: alg.NewVector3(0, 0, 1), t1: -1, t2: 1},
	}
	for _, e := range ex {
		r := NewRay(e.p, e.d)
		xs := c.Intersect(r)
		AssertTrue(t, len(xs) == 2)
		AssertFloatEqual(t, xs[0].Distance, e.t1)
		AssertFloatEqual(t, xs[1].Distance, e.t2)
	}
}

func TestARayMissesACube(t *testing.T) {
	c := NewCube()
	ex := []struct {
		p alg.Vector
		d alg.Vector
	}{
		{p: alg.NewPoint(-2, 0, 0), d: alg.NewVector3(0.2673, 0.5345, 0.8018)},
		{p: alg.NewPoint(0, -2, 0), d: alg.NewVector3(0.8018, 0.2673, 0.5345)},
		{p: alg.NewPoint(0, 0, -2), d: alg.NewVector3(0.5345, 0.8018, 0.2673)},
		{p: alg.NewPoint(2, 0, 2), d: alg.NewVector3(0, 0, -1)},
		{p: alg.NewPoint(0, 2, 2), d: alg.NewVector3(0, -1, 0)},
		{p: alg.NewPoint(2, 2, 0), d: alg.NewVector3(-1, 0, 0)},
	}
	for _, e := range ex {
		r := NewRay(e.p, e.d)
		xs := c.Intersect(r)
		AssertTrue(t, len(xs) == 0)
	}
}

func TestTheNormalOnTheSurfaceOfACube(t *testing.T) {
	c := NewCube()
	ex := []struct {
		p alg.Vector
		n alg.Vector
	}{
		{p: alg.NewPoint(1, 0.5, -0.8), n: alg.NewVector3(1, 0, 0)},
		{p: alg.NewPoint(-1, -0.2, 0.9), n: alg.NewVector3(-1, 0, 0)},
		{p: alg.NewPoint(-0.4, 1, -0.1), n: alg.NewVector3(0, 1, 0)},
		{p: alg.NewPoint(0.3, -1, -0.7), n: alg.NewVector3(0, -1, 0)},
		{p: alg.NewPoint(-0.6, 0.3, 1), n: alg.NewVector3(0, 0, 1)},
		{p: alg.NewPoint(0.4, 0., -1), n: alg.NewVector3(0, 0, -1)},
		{p: alg.NewPoint(1, 1, 1), n: alg.NewVector3(1, 0, 0)},
		{p: alg.NewPoint(-1, -1, -1), n: alg.NewVector3(-1, 0, 0)},
	}
	for _, e := range ex {
		AssertVectorEqual(t, c.NormalAt(e.p, nil), e.n)
	}
}
