package render

import (
	"sort"

	"github.com/mhoertnagl/gracer/alg"
)

type Object interface {
	Intersect(r *Ray) Intersections
	NormalAt(p alg.Vector, hit *Intersection) alg.Vector
	GetMaterial() *Material
	GetTransform() alg.Matrix
	GetInverseTransform() alg.Matrix
	SetParent(obj Object)
	GetParent() Object
	GetBounds() *Bounds
}

// TODO: Add another interface Shape that inherits from Object with GetMaterial().

func worldToObject(obj Object, point alg.Vector) alg.Vector {
	if parent := obj.GetParent(); parent != nil {
		point = worldToObject(parent, point)
	}
	return obj.GetInverseTransform().MultVec(point)
}

func normalToWorld(obj Object, normal alg.Vector) alg.Vector {
	normal = obj.GetInverseTransform().Transpose().MultVec(normal)
	// Reset w coordinate to 0.
	normal[3] = 0
	normal = normal.Norm()
	if parent := obj.GetParent(); parent != nil {
		normal = normalToWorld(parent, normal)
	}
	return normal
}

// TODO: Custom type Objects []Object with (objs Objects) Intersect(r *Ray)
func IntersectCollection(objs []Object, r *Ray) Intersections {
	xxs := NewIntersections()
	for _, obj := range objs {
		xs := obj.Intersect(r)
		xxs = append(xxs, xs...)
	}
	sort.Slice(xxs, func(i, j int) bool {
		return xxs[i].Distance < xxs[j].Distance
	})
	return xxs
}
