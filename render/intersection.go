package render

type Intersection struct {
	t   float64
	obj Object
}

func NewIntersection(t float64, obj Object) *Intersection {
	return &Intersection{t: t, obj: obj}
}

type Intersections []*Intersection

func NewIntersections(is ...*Intersection) Intersections {
	return is
}

func (xs Intersections) Hit() *Intersection {
	var hit *Intersection = nil
	for i := 0; i < len(xs); i++ {
		if x := xs[i]; x.t > 0 && (hit == nil || x.t < hit.t) {
			hit = xs[i]
		}
	}
	return hit
}
