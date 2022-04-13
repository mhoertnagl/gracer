package render

import "github.com/mhoertnagl/gracer/alg"

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
	// or := r.Transform(alg.Inverse(t.Transform))
	return NewIntersections()
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
	return NewBounds(alg.NewPoint(-1, -1, -1), alg.NewPoint(1, 1, 1))
}
