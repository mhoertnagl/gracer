package render

import (
	"math"
	"sync"

	"github.com/mhoertnagl/gracer/canvas"
)

const EPSILON = 2.220446049250313e-6

// TODO: Use a Group as Objects root.
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

type line struct {
	y  int
	xs []canvas.Color
}

func newLine(y int, len int) *line {
	return &line{
		y:  y,
		xs: make([]canvas.Color, len),
	}
}

func (world *World) Render(camera *Camera) *canvas.Canvas {
	canvas := canvas.NewCanvas(camera.hsize, camera.vsize)
	ch := make(chan *line, camera.vsize)
	wg := sync.WaitGroup{}
	for y := 0; y < camera.vsize; y++ {
		wg.Add(1)
		go renderLine(world, camera, y, ch, &wg)
	}
	wg.Wait()
	close(ch)
	for line := range ch {
		for x := 0; x < camera.hsize; x++ {
			canvas.Set(x, line.y, line.xs[x])
		}
	}
	return canvas
}

func renderLine(world *World, camera *Camera, y int, ch chan *line, wg *sync.WaitGroup) {
	line := newLine(y, camera.hsize)
	for x := 0; x < camera.hsize; x++ {
		ray := camera.RayForPixel(x, y)
		color := world.colorAt(ray, world.MaxBounces)
		line.xs[x] = color
	}
	ch <- line
	wg.Done()
}

func (w *World) colorAt(r *Ray, remaining int) canvas.Color {
	xs := w.intersect(r)
	if hit := xs.Hit(); hit != nil {
		c := prepareComps(hit, r, xs)
		return w.shade(c, remaining)
	}
	return canvas.Black
}

// TODO: Same as group.intersect
func (w *World) intersect(r *Ray) Intersections {
	return IntersectCollection(w.Objects, r)
}

func (w *World) shade(c *comps, remaining int) canvas.Color {
	color := canvas.Black
	for _, light := range w.Lights {
		isShadowed := false
		if c.Object.GetMaterial().ReceiveShadow {
			isShadowed = light.IsShadowed(w, c.OverPoint)
		}
		surface := light.Lighting(c.Object, c.OverPoint, c.Eye, c.Normal, isShadowed)
		color = color.Add(surface)
	}
	reflected := w.reflectedColor(c, remaining)
	refracted := w.refractedColor(c, remaining)
	material := c.Object.GetMaterial()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := schlick(c)
		return color.Add(reflected.Scale(reflectance)).Add(refracted.Scale(1.0 - reflectance))
	}
	return color.Add(reflected).Add(refracted)
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

func (w *World) refractedColor(c *comps, remaining int) canvas.Color {
	if remaining <= 0 {
		return canvas.Black
	}
	material := c.Object.GetMaterial()
	if material.Transparency == 0 {
		return canvas.Black
	}
	nratio := c.N1 / c.N2
	cosi := c.Eye.Dot(c.Normal)
	sin2t := nratio * nratio * (1.0 - cosi*cosi)
	if sin2t > 1 {
		return canvas.Black
	}
	cost := math.Sqrt(1.0 - sin2t)
	direction := c.Normal.Mult(nratio*cosi - cost).Sub(c.Eye.Mult(nratio))
	refractray := NewRay(c.UnderPoint, direction)
	return w.colorAt(refractray, remaining-1).Scale(material.Transparency)
}
