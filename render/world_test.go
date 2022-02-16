package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

func TestIntersectDefaultWorldWithARay(t *testing.T) {
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	xs := w.intersect(r)
	AssertIntEqual(t, len(xs), 4)
	AssertFloatEqual(t, xs[0].Distance, 4)
	AssertFloatEqual(t, xs[1].Distance, 4.5)
	AssertFloatEqual(t, xs[2].Distance, 5.5)
	AssertFloatEqual(t, xs[3].Distance, 6)
}

func TestPrecomputingTheStateOfAnIntersection(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := prepareComps(i, r)
	AssertFloatEqual(t, c.Distance, i.Distance)
	AssertObjectEqual(t, c.Object, i.Object)
	AssertVectorEqual(t, c.Point, alg.NewPoint(0, 0, -1))
	AssertVectorEqual(t, c.Eye, alg.NewVector3(0, 0, -1))
	AssertVectorEqual(t, c.Normal, alg.NewVector3(0, 0, -1))
}

func TestTheHitWhenAnIntersectionOccursOnTheOutside(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := prepareComps(i, r)
	AssertFalse(t, c.Inside)
}

func TestTheHitWhenAnIntersectionOccursOnTheInside(t *testing.T) {
	r := NewRay(alg.NewPoint(0, 0, 0), alg.NewVector3(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(1, s)
	c := prepareComps(i, r)
	AssertTrue(t, c.Inside)
	AssertVectorEqual(t, c.Point, alg.NewPoint(0, 0, 1))
	AssertVectorEqual(t, c.Eye, alg.NewVector3(0, 0, -1))
	AssertVectorEqual(t, c.Normal, alg.NewVector3(0, 0, -1))
}

func TestShadingAnIntersection(t *testing.T) {
	e := canvas.NewColor(0.38066, 0.47583, 0.28550)
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	s := w.Objects[0]
	i := NewIntersection(4, s)
	c := prepareComps(i, r)
	AssertColorEqual(t, w.shade(c), e)
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	e := canvas.NewColor(0.90498, 0.90498, 0.90498)
	w := newDefaultWorld()
	w.ClearLights()
	l := NewPointLight(alg.NewPoint(0, 0.25, 0), canvas.White())
	w.AddLight(l)
	r := NewRay(alg.NewPoint(0, 0, 0), alg.NewVector3(0, 0, 1))
	s := w.Objects[1]
	i := NewIntersection(0.5, s)
	c := prepareComps(i, r)
	AssertColorEqual(t, w.shade(c), e)
}

func TestTheColorWhenTheRayMisses(t *testing.T) {
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 1, 0))
	c := w.ColorAt(r)
	AssertColorEqual(t, c, canvas.Black())
}

func TestTheColorWhenTheRayHits(t *testing.T) {
	e := canvas.NewColor(0.38066, 0.47583, 0.2855)
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	c := w.ColorAt(r)
	AssertColorEqual(t, c, e)
}

func TestTheColorWithAnIntersectionBehindTheRay(t *testing.T) {
	w := newDefaultWorld()
	w.Objects[0].GetMaterial().Ambient = 1
	w.Objects[1].GetMaterial().Ambient = 1
	r := NewRay(alg.NewPoint(0, 0, 0.75), alg.NewVector3(0, 0, -1))
	c := w.ColorAt(r)
	AssertColorEqual(t, c, w.Objects[1].GetMaterial().Color)
}

func newDefaultWorld() *World {
	w := NewWorld()
	lp := alg.NewPoint(-10, 10, -10)
	l := NewPointLight(lp, canvas.White())
	s1 := NewSphere()
	s1.Material.Color = canvas.NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2 := NewSphere()
	s2.Transform = alg.Scaling(0.5, 0.5, 0.5)
	w.AddLight(l)
	w.AddObject(s1)
	w.AddObject(s2)
	return w
}
