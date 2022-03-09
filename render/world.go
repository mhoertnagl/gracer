package render

import (
	"sort"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

const EPSILON = 2.220446049250313e-8

type World struct {
	Lights     []Light
	Objects    []Object
	MaxBounces int
}

func NewWorld() *World {
	return &World{
		Lights:     make([]Light, 0),
		Objects:    make([]Object, 0),
		MaxBounces: 5,
	}
}

func (w *World) AddLight(light Light) {
	w.Lights = append(w.Lights, light)
}

func (w *World) AddObject(object Object) {
	w.Objects = append(w.Objects, object)
}

func (world *World) Render(camera *Camera) *canvas.Canvas {
	canvas := canvas.NewCanvas(camera.hsize, camera.vsize)
	for y := 0; y < camera.vsize; y++ {
		for x := 0; x < camera.hsize; x++ {
			ray := camera.RayForPixel(x, y)
			color := world.colorAt(ray, world.MaxBounces)
			canvas.Set(x, y, color)
		}
	}
	return canvas
}

func (w *World) colorAt(r *Ray, remaining int) canvas.Color {
	xs := w.intersect(r)
	if hit := xs.Hit(); hit != nil {
		c := prepareComps(hit, r)
		return w.shade(c, remaining)
	}
	return canvas.Black
}

func (w *World) intersect(r *Ray) Intersections {
	xxs := Intersections{}
	for _, object := range w.Objects {
		xs := object.Intersect(r)
		xxs = append(xxs, xs...)
	}
	sort.Slice(xxs, func(i, j int) bool {
		return xxs[i].Distance < xxs[j].Distance
	})
	return xxs
}

// TODO: Why isn't this part of the intersection?
type comps struct {
	Distance  float64
	Object    Object
	Point     alg.Vector
	OverPoint alg.Vector
	Eye       alg.Vector
	Normal    alg.Vector
	Reflect   alg.Vector
	Inside    bool
}

func prepareComps(i *Intersection, r *Ray) *comps {
	c := &comps{}
	c.Distance = i.Distance
	c.Object = i.Object
	c.Point = r.Position(i.Distance)
	c.Eye = r.Direction.Neg()
	c.Normal = i.Object.NormalAt(c.Point)
	if c.Normal.Dot(c.Eye) < 0 {
		c.Inside = true
		c.Normal = c.Normal.Neg()
	}
	c.OverPoint = c.Point.Add(c.Normal.Mult(EPSILON))
	c.Reflect = r.Direction.Reflect(c.Normal)
	return c
}

func (w *World) shade(c *comps, remaining int) canvas.Color {
	color := canvas.Black
	for _, light := range w.Lights {
		isShadowed := light.IsShadowed(w, c.OverPoint)
		surface := light.Lighting(c.Object, c.OverPoint, c.Eye, c.Normal, isShadowed)
		color = color.Add(surface)
	}
	reflected := w.reflectedColor(c, remaining)
	return color.Add(reflected)
}

func (w *World) reflectedColor(c *comps, remaining int) canvas.Color {
	if remaining <= 0 {
		return canvas.Black
	}
	material := c.Object.GetMaterial()
	if material.Reflective == 0 {
		return canvas.Black
	}
	r := NewRay(c.OverPoint, c.Reflect)
	color := w.colorAt(r, remaining-1)
	return color.Scale(material.Reflective)
}
