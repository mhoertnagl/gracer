package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type SmoothTriangle struct {
	P1           alg.Vector
	P2           alg.Vector
	P3           alg.Vector
	E1           alg.Vector
	E2           alg.Vector
	N1           alg.Vector
	N2           alg.Vector
	N3           alg.Vector
	Transform    alg.Matrix
	Material     *Material
	Parent       Object
	invTransform alg.Matrix
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 alg.Vector) *SmoothTriangle {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)
	return &SmoothTriangle{
		P1:           p1,
		P2:           p2,
		P3:           p3,
		E1:           e1,
		E2:           e2,
		N1:           n1,
		N2:           n2,
		N3:           n3,
		Transform:    alg.Id4,
		Material:     NewMaterial(),
		Parent:       nil,
		invTransform: nil,
	}
}

func (t *SmoothTriangle) Intersect(r *Ray) Intersections {
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
		NewIntersectionWithUV(w, t, u, v),
	)
}

func (t *SmoothTriangle) NormalAt(point alg.Vector, hit *Intersection) alg.Vector {
	op := worldToObject(t, point)
	on := t.localNormalAt(op, hit)
	return normalToWorld(t, on)
}

func (t *SmoothTriangle) localNormalAt(point alg.Vector, hit *Intersection) alg.Vector {
	n1 := t.N1.Mult(1 - hit.U - hit.V)
	n2 := t.N2.Mult(hit.U)
	n3 := t.N3.Mult(hit.V)
	return n1.Add(n2).Add(n3)
}

func (t *SmoothTriangle) GetMaterial() *Material {
	return t.Material
}

func (t *SmoothTriangle) GetTransform() alg.Matrix {
	return t.Transform
}

func (t *SmoothTriangle) SetParent(obj Object) {
	t.Parent = obj
}

func (t *SmoothTriangle) GetParent() Object {
	return t.Parent
}

func (t *SmoothTriangle) GetInverseTransform() alg.Matrix {
	if t.invTransform == nil {
		t.invTransform = alg.Inverse(t.Transform)
	}
	return t.invTransform
}

func (t *SmoothTriangle) GetBounds() *Bounds {
	min := alg.Min3(t.P1, t.P2, t.P3)
	max := alg.Max3(t.P1, t.P2, t.P3)
	return NewBounds(min, max)
}
