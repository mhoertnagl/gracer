package render

import (
	"github.com/mhoertnagl/gracer/alg"
)

type Group struct {
	Transform alg.Matrix
	Parent    Object
	Kids      []Object
}

func NewGroup() *Group {
	return &Group{
		Transform: alg.Id4,
		Parent:    nil,
		Kids:      make([]Object, 0),
	}
}

func (g *Group) Intersect(r *Ray) Intersections {
	r2 := r.Transform(alg.Inverse(g.Transform))
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
