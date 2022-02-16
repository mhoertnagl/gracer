package render

import "sort"

type World struct {
	Lights  []Light
	Objects []Object
}

func NewWorld() *World {
	return &World{
		Lights:  make([]Light, 0),
		Objects: make([]Object, 0),
	}
}

func (w *World) AddLight(light Light) {
	w.Lights = append(w.Lights, light)
}

func (w *World) AddObject(object Object) {
	w.Objects = append(w.Objects, object)
}

func (w *World) Intersect(r *Ray) Intersections {
	xxs := Intersections{}
	for _, obj := range w.Objects {
		xs := obj.Intersect(r)
		xxs = append(xxs, xs...)
	}
	sort.Slice(xxs, func(i, j int) bool {
		return xxs[i].Distance < xxs[j].Distance
	})
	return xxs
}
