package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Plane struct {
	Transform alg.Matrix
	Material  *Material
}

func NewPlane() *Plane {
	return &Plane{
		Transform: alg.Id4,
		Material:  NewMaterial(),
	}
}

func (p *Plane) Intersect(r *Ray) Intersections {
	or := r.Transform(alg.Inverse(p.Transform))
	if math.Abs(or.Direction[1]) < EPSILON {
		return Intersections{}
	}
	return Intersections{
		NewIntersection(-or.Origin[1]/or.Direction[1], p),
	}
}

func (p *Plane) NormalAt(point alg.Vector) alg.Vector {
	inv := alg.Inverse(p.Transform)
	on := alg.NewVector3(0, 1, 0)
	wn := inv.Transpose().MultVec(on)
	// Reset w coordinate to 0.
	wn[3] = 0
	return wn.Norm()
}

func (p *Plane) GetMaterial() *Material {
	return p.Material
}

func (p *Plane) GetTransform() alg.Matrix {
	return p.Transform
}
