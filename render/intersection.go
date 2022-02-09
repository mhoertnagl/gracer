package render

type Intersection struct {
	Distance float64
	Obj      Object
}

func NewIntersection(distance float64, obj Object) *Intersection {
	return &Intersection{Distance: distance, Obj: obj}
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
