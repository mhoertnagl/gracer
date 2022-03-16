package render

import (
	"container/list"

	"github.com/mhoertnagl/gracer/alg"
)

type Intersection struct {
	Distance float64
	Object   Object
}

func NewIntersection(distance float64, object Object) *Intersection {
	return &Intersection{Distance: distance, Object: object}
}

type Intersections []*Intersection

func NewIntersections(is ...*Intersection) Intersections {
	return is
}

func (xs Intersections) Hit() *Intersection {
	var hit *Intersection = nil
	for i := 0; i < len(xs); i++ {
		if x := xs[i]; x.Distance > 0 && (hit == nil || x.Distance < hit.Distance) {
			hit = xs[i]
		}
	}
	return hit
}

// TODO: Why isn't this part of the intersection?
type comps struct {
	Distance   float64
	Object     Object
	Point      alg.Vector
	OverPoint  alg.Vector
	UnderPoint alg.Vector
	Eye        alg.Vector
	Normal     alg.Vector
	Reflect    alg.Vector
	N1         float64
	N2         float64
	Inside     bool
}

func prepareComps(hit *Intersection, r *Ray, xs Intersections) *comps {
	c := &comps{}
	c.Distance = hit.Distance
	c.Object = hit.Object
	c.Point = r.Position(hit.Distance)
	c.Eye = r.Direction.Neg()
	c.Normal = hit.Object.NormalAt(c.Point)
	if c.Normal.Dot(c.Eye) < 0 {
		c.Inside = true
		c.Normal = c.Normal.Neg()
	}
	delta := c.Normal.Mult(EPSILON)
	c.OverPoint = c.Point.Add(delta)
	c.UnderPoint = c.Point.Sub(delta)
	c.Reflect = r.Direction.Reflect(c.Normal)

	containers := list.New()
	for _, x := range xs {
		if x == hit {
			if listEmpty(containers) {
				c.N1 = 1.0
			} else {
				elem := containers.Back()
				obj := elem.Value.(Object)
				c.N1 = obj.GetMaterial().RefractiveIndex
			}
		}

		if elem := listFind(containers, x.Object); elem != nil {
			containers.Remove(elem)
		} else {
			containers.PushBack(x.Object)
			// containers.PushFront(x.Object)
		}

		if x == hit {
			if listEmpty(containers) {
				c.N2 = 1.0
			} else {
				elem := containers.Back()
				obj := elem.Value.(Object)
				c.N2 = obj.GetMaterial().RefractiveIndex
			}
		}
	}
	return c
}

func listEmpty(items *list.List) bool {
	return items.Front() == nil
}

func listFind(items *list.List, item interface{}) *list.Element {
	for elem := items.Back(); elem != nil; elem = elem.Prev() {
		if elem.Value == item {
			return elem
		}
	}
	return nil
}
