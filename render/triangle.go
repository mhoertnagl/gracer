package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Triangle struct {
	P1           alg.Vector
	P2           alg.Vector
	P3           alg.Vector
	E1           alg.Vector
	E2           alg.Vector
	Normal       alg.Vector
	Transform    alg.Matrix
	Material     *Material
	Parent       Object
	invTransform alg.Matrix
}

func NewTriangle(p1, p2, p3 alg.Vector) *Triangle {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)
	n := e2.Cross(e1).Norm()
	return &Triangle{
		P1:           p1,
		P2:           p2,
		P3:           p3,
		E1:           e1,
		E2:           e2,
		Normal:       n,
		Transform:    alg.Id4,
		Material:     NewMaterial(),
		Parent:       nil,
		invTransform: nil,
	}
}

func (t *Triangle) Intersect(r *Ray) Intersections {
	dirCrossE2 := r.Direction.Cross(t.E2)
	det := t.E1.Dot(dirCrossE2)
	if math.Abs(det) < EPSILON {
		return NewIntersections()
	}
	f := 1.0 / det
	p1ToOrigin := r.Origin.Sub(t.P1)
	u := f * p1ToOrigin.Dot(dirCrossE2)
	if u < 0 || u > 1 {
		return NewIntersections()
	}
	origCrossE1 := p1ToOrigin.Cross(t.E1)
	v := f * r.Direction.Dot(origCrossE1)
	if v < 0 || (u+v) > 1 {
		return NewIntersections()
	}
	w := f * t.E2.Dot(origCrossE1)
	return NewIntersections(
		NewIntersection(w, t),
	)
}

func (t *Triangle) NormalAt(point alg.Vector) alg.Vector {
	op := worldToObject(t, point)
	on := t.localNormalAt(op)
	return normalToWorld(t, on)
}

func (t *Triangle) localNormalAt(point alg.Vector) alg.Vector {
	return t.Normal
}

func (t *Triangle) GetMaterial() *Material {
	return t.Material
}

func (t *Triangle) GetTransform() alg.Matrix {
	return t.Transform
}

func (t *Triangle) SetParent(obj Object) {
	t.Parent = obj
}

func (t *Triangle) GetParent() Object {
	return t.Parent
}

func (t *Triangle) GetInverseTransform() alg.Matrix {
	if t.invTransform == nil {
		t.invTransform = alg.Inverse(t.Transform)
	}
	return t.invTransform
}

func (t *Triangle) GetBounds() *Bounds {
	min := alg.Min3(t.P1, t.P2, t.P3)
	max := alg.Max3(t.P1, t.P2, t.P3)
	return NewBounds(min, max)
}
