package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Sphere struct {
	Transform alg.Matrix
	Material  *Material
	Parent    Object
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: alg.Id4,
		Material:  NewMaterial(),
		Parent:    nil,
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

func (s *Sphere) NormalAt(point alg.Vector) alg.Vector {
	op := worldToObject(s, point)
	on := s.localNormalAt(op)
	return normalToWorld(s, on)
	// inv := alg.Inverse(s.Transform)
	// op := inv.MultVec(p)
	// on := op.Sub(alg.Origin)
	// wn := inv.Transpose().MultVec(on)
	// // Reset w coordinate to 0.
	// wn[3] = 0
	// return wn.Norm()
}

func (s *Sphere) localNormalAt(point alg.Vector) alg.Vector {
	return point.Sub(alg.Origin)
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}

func (s *Sphere) GetTransform() alg.Matrix {
	return s.Transform
}

func (s *Sphere) SetParent(obj Object) {
	s.Parent = obj
}

func (s *Sphere) GetParent() Object {
	return s.Parent
}
