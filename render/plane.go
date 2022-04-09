package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Plane struct {
	Transform alg.Matrix
	Material  *Material
	Parent    Object
}

func NewPlane() *Plane {
	return &Plane{
		Transform: alg.Id4,
		Material:  NewMaterial(),
		Parent:    nil,
	}
}

func (p *Plane) Intersect(r *Ray) Intersections {
	or := r.Transform(alg.Inverse(p.Transform))
	if math.Abs(or.Direction[1]) < EPSILON {
		return NewIntersections()
	}
	return NewIntersections(
		NewIntersection(-or.Origin[1]/or.Direction[1], p),
	)
}

func (p *Plane) NormalAt(point alg.Vector) alg.Vector {
	op := worldToObject(p, point)
	on := p.localNormalAt(op)
	return normalToWorld(p, on)
}

func (p *Plane) localNormalAt(point alg.Vector) alg.Vector {
	return alg.NewVector3(0, 1, 0)
}

func (p *Plane) GetMaterial() *Material {
	return p.Material
}

func (p *Plane) GetTransform() alg.Matrix {
	return p.Transform
}

func (p *Plane) SetParent(obj Object) {
	p.Parent = obj
}

func (p *Plane) GetParent() Object {
	return p.Parent
}
