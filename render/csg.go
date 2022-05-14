package render

import (
	"github.com/mhoertnagl/gracer/alg"
)

type CsgOp int

const (
	Union CsgOp = iota
	Intersect
	Difference
)

type Csg struct {
	Transform    alg.Matrix
	Parent       Object
	Op           CsgOp
	Left         Object
	Right        Object
	invTransform alg.Matrix
	bounds       *Bounds
}

func NewUnion(left, right Object) *Csg {
	return &Csg{
		Transform:    alg.Id4,
		Parent:       nil,
		Op:           Union,
		Left:         left,
		Right:        right,
		invTransform: nil,
		bounds:       nil,
	}
}

func NewIntersect(left, right Object) *Csg {
	return &Csg{
		Transform:    alg.Id4,
		Parent:       nil,
		Op:           Intersect,
		Left:         left,
		Right:        right,
		invTransform: nil,
		bounds:       nil,
	}
}

func NewDifference(left, right Object) *Csg {
	return &Csg{
		Transform:    alg.Id4,
		Parent:       nil,
		Op:           Difference,
		Left:         left,
		Right:        right,
		invTransform: nil,
		bounds:       nil,
	}
}

func (g *Csg) Intersect(r *Ray) Intersections {
	or := r.Transform(g.GetInverseTransform())
	xs := IntersectCollection([]Object{g.Left, g.Right}, or)
	return filter_intersections(g, xs)
}

func (g *Csg) NormalAt(p alg.Vector, hit *Intersection) alg.Vector {
	panic("  don't have normal vectors")
}

func (g *Csg) GetMaterial() *Material {
	panic("Groups don't have materials")
}

func (g *Csg) GetTransform() alg.Matrix {
	return g.Transform
}

func (g *Csg) SetParent(obj Object) {
	g.Parent = obj
}

func (g *Csg) GetParent() Object {
	return g.Parent
}

func (g *Csg) SetLeft(kid Object) {
	kid.SetParent(g)
	g.Left = g
}

func (g *Csg) GetLeft() Object {
	return g.Left
}

func (g *Csg) SetRight(kid Object) {
	kid.SetParent(g)
	g.Right = g
}

func (g *Csg) GetRight() Object {
	return g.Right
}

func (g *Csg) GetInverseTransform() alg.Matrix {
	if g.invTransform == nil {
		g.invTransform = alg.Inverse(g.Transform)
	}
	return g.invTransform
}

func (g *Csg) GetBounds() *Bounds {
	if g.bounds == nil {
		edges := []alg.Vector{}
		lbe := g.Left.GetBounds().Edges()
		rbe := g.Right.GetBounds().Edges()
		edges = append(edges, lbe...)
		edges = append(edges, rbe...)
		g.bounds = NewBoundsFrom(edges...)
	}
	return g.bounds
}

func (g *Csg) Includes(obj Object) bool {
	return g.Left.Includes(obj) || g.Right.Includes(obj)
}

func intersectionAllowed(op CsgOp, lhit, inl, inr bool) bool {
	switch op {
	case Union:
		return (lhit && !inr) || (!lhit && !inl)
	case Intersect:
		return (lhit && inr) || (!lhit && inl)
	case Difference:
		return (lhit && !inr) || (!lhit && inl)
	}
	return false
}

func filter_intersections(csg *Csg, xs Intersections) Intersections {
	inl := false
	inr := false
	res := NewIntersections()
	for _, x := range xs {
		lhit := csg.Left.Includes(x.Object)
		if intersectionAllowed(csg.Op, lhit, inl, inr) {
			res = append(res, x)
		}
		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}
	return res
}
