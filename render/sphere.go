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
	or := r.Transform(alg.Inverse(s.Transform))
	str := or.Origin.Sub(alg.Origin)
	a := or.Direction.Dot(or.Direction)
	b := 2 * or.Direction.Dot(str)
	c := str.Dot(str) - 1
	d := b*b - 4*a*c
	if d < 0 {
		return NewIntersections()
	}
	ds := math.Sqrt(d)
	a2 := 2 * a
	return NewIntersections(
		NewIntersection((-b-ds)/a2, s),
		NewIntersection((-b+ds)/a2, s),
	)
}

func (s *Sphere) NormalAt(point alg.Vector) alg.Vector {
	op := worldToObject(s, point)
	on := s.localNormalAt(op)
	return normalToWorld(s, on)
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

func (s *Sphere) Bounds() *Bounds {
	return NewBounds(alg.NewPoint(-1, -1, -1), alg.NewPoint(1, 1, 1))
}
