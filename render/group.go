package render

import (
	"github.com/mhoertnagl/gracer/alg"
)

type Group struct {
	Transform    alg.Matrix
	Parent       Object
	Kids         []Object
	invTransform alg.Matrix
	bounds       *Bounds
}

func NewGroup() *Group {
	return &Group{
		Transform:    alg.Id4,
		Parent:       nil,
		Kids:         make([]Object, 0),
		invTransform: nil,
		bounds:       nil,
	}
}

func (g *Group) Intersect(r *Ray) Intersections {
	r2 := r.Transform(g.GetInverseTransform())
	return IntersectCollection(g.Kids, r2)
}

func (g *Group) NormalAt(p alg.Vector) alg.Vector {
	return alg.NewZeroVector(4)
}

func (g *Group) GetMaterial() *Material {
	panic("Groups don't have materials")
}

func (g *Group) GetTransform() alg.Matrix {
	return g.Transform
}

func (g *Group) SetParent(obj Object) {
	g.Parent = obj
}

func (g *Group) GetParent() Object {
	return g.Parent
}

func (g *Group) AddKid(kid Object) {
	kid.SetParent(g)
	g.Kids = append(g.Kids, kid)
}

func (g *Group) GetInverseTransform() alg.Matrix {
	if g.invTransform == nil {
		g.invTransform = alg.Inverse(g.Transform)
	}
	return g.invTransform
}

func (g *Group) GetBounds() *Bounds {
	if g.bounds == nil {
		g.bounds = NewBounds(alg.NewPoint(-1, -1, -1), alg.NewPoint(1, 1, 1))
	}
	return g.bounds
}
