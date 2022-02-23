package render

import (
	"sort"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

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

func (world *World) Render(camera *Camera) *canvas.Canvas {
	canvas := canvas.NewCanvas(camera.hsize, camera.vsize)
	for y := 0; y < camera.vsize; y++ {
		for x := 0; x < camera.hsize; x++ {
			ray := camera.RayForPixel(x, y)
			color := world.colorAt(ray)
			canvas.Set(x, y, color)
		}
	}
	return canvas
}

func (w *World) colorAt(r *Ray) canvas.Color {
	xs := w.intersect(r)
	if hit := xs.Hit(); hit != nil {
		c := prepareComps(hit, r)
		return w.shade(c)
	}
	return canvas.Black()
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
	Distance float64
	Object   Object
	Point    alg.Vector
	Eye      alg.Vector
	Normal   alg.Vector
	Inside   bool
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
	return c
}

func (w *World) shade(c *comps) canvas.Color {
	color := canvas.Black()
	for _, light := range w.Lights {
		m := c.Object.GetMaterial()
		c := light.Lighting(m, c.Point, c.Eye, c.Normal)
		color = color.Add(c)
	}
	return color
}
