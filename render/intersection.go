package render

import (
	"container/list"
	"math"
	"reflect"

	"github.com/mhoertnagl/gracer/alg"
)

type Intersection struct {
	Distance float64
	Object   Object
	U        float64
	V        float64
}

func NewIntersection(distance float64, object Object) *Intersection {
	return &Intersection{Distance: distance, Object: object}
}

func NewIntersectionWithUV(distance float64, object Object, u, v float64) *Intersection {
	return &Intersection{Distance: distance, Object: object, U: u, V: v}
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
	c.Normal = hit.Object.NormalAt(c.Point, hit)
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
		if reflect.DeepEqual(elem.Value, item) {
			return elem
		}
	}
	return nil
}

func schlick(c *comps) float64 {
	cos := c.Eye.Dot(c.Normal)
	if c.N1 > c.N2 {
		N := c.N1 / c.N2
		sin2t := N * N * (1.0 - cos*cos)
		if sin2t > 1.0 {
			return 1.0
		}
		cos = math.Sqrt(1.0 - sin2t)
	}
	k0 := ((c.N1 - c.N2) / (c.N1 + c.N2))
	r0 := k0 * k0
	return r0 + (1.0-r0)*math.Pow((1.0-cos), 5)
}
