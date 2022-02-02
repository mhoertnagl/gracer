package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Object interface {
	Intersect(r Ray) Intersections
}

type Sphere struct {
	Transform alg.Matrix
}

func NewSphere() Sphere {
	return Sphere{Transform: alg.Id4}
}

func (s Sphere) Intersect(r Ray) Intersections {
	r2 := r.Transform(alg.Inverse(s.Transform))
	str := r2.origin.Sub(alg.NewPoint(0, 0, 0))
	a := r2.direction.Dot(r2.direction)
	b := 2 * r2.direction.Dot(str)
	c := str.Dot(str) - 1
	d := b*b - 4*a*c
	if d < 0 {
		return Intersections{}
	}
	ds := math.Sqrt(d)
	a2 := 2 * a
	return Intersections{
		NewIntersection((-b-ds)/a2, s),
		NewIntersection((-b+ds)/a2, s),
	}
}
