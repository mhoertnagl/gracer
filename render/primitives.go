package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Object interface {
	Intersect(r *Ray) Intersections
	NormalAt(p alg.Vector) alg.Vector
	GetMaterial() *Material
}

type Sphere struct {
	Transform alg.Matrix
	Material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: alg.Id4,
		Material:  NewMaterial(),
	}
}

func (s *Sphere) Intersect(r *Ray) Intersections {
	r2 := r.Transform(alg.Inverse(s.Transform))
	str := r2.Origin.Sub(alg.Origin)
	a := r2.Direction.Dot(r2.Direction)
	b := 2 * r2.Direction.Dot(str)
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

func (s *Sphere) NormalAt(p alg.Vector) alg.Vector {
	inv := alg.Inverse(s.Transform)
	op := inv.MultVec(p)
	on := op.Sub(alg.Origin)
	n := inv.Transpose().MultVec(on)
	// Reset w coordinate to 0.
	n[3] = 0
	return n.Norm()
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}
