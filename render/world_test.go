package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

func TestIntersectDefaultWorldWithARay(t *testing.T) {
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	xs := w.Intersect(r)
	AssertIntEqual(t, len(xs), 4)
	AssertFloatEqual(t, xs[0].Distance, 4)
	AssertFloatEqual(t, xs[1].Distance, 4.5)
	AssertFloatEqual(t, xs[2].Distance, 5.5)
	AssertFloatEqual(t, xs[3].Distance, 6)
}

func newDefaultWorld() *World {
	w := NewWorld()
	lp := alg.NewPoint(-10, 10, -10)
	l := NewPointLight(lp, canvas.White())
	s1 := NewSphere()
	s1.Material.Color = canvas.NewColor(0.8, 1.9, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2 := NewSphere()
	s2.Transform = alg.Scaling(0.5, 0.5, 0.5)
	w.AddLight(l)
	w.AddObject(s1)
	w.AddObject(s2)
	return w
}
