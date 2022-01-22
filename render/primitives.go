package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Object interface {
	Intersect(r Ray) Intersections
}

type Sphere struct{}

func NewSphere() Sphere {
	return Sphere{}
}

func (s Sphere) Intersect(r Ray) Intersections {
	str := r.origin.Sub(alg.NewPoint(0, 0, 0))
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(str)
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
