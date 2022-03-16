package render

import (
	"math"
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

func TestTheColorWhenTheRayMisses(t *testing.T) {
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 1, 0))
	c := w.colorAt(r, w.MaxBounces)
	AssertColorEqual(t, c, canvas.Black)
}

func TestTheColorWhenTheRayHits(t *testing.T) {
	e := canvas.NewColor(0.38066, 0.47583, 0.2855)
	w := newDefaultWorld()
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	c := w.colorAt(r, w.MaxBounces)
	AssertColorEqual(t, c, e)
}

func TestTheColorWithAnIntersectionBehindTheRay(t *testing.T) {
	w := newDefaultWorld()
	w.Objects[0].GetMaterial().Ambient = 1
	w.Objects[1].GetMaterial().Ambient = 1
	r := NewRay(alg.NewPoint(0, 0, 0.75), alg.NewVector3(0, 0, -1))
	c := w.colorAt(r, w.MaxBounces)
	AssertColorEqual(t, c, w.Objects[1].GetMaterial().Color)
}

func TestTheReflectedColorForANonreflectiveMaterial(t *testing.T) {
	w := newDefaultWorld()
	w.Objects[1].GetMaterial().Ambient = 1
	r := NewRay(alg.NewPoint(0, 0, 0), alg.NewVector3(0, 0, 1))
	i := NewIntersection(1, w.Objects[1])
	c := prepareComps(i, r, NewIntersections(i))
	color := w.reflectedColor(c, w.MaxBounces)
	AssertColorEqual(t, color, canvas.Black)
}

func TestTheReflectedColorForAReflectiveMaterial(t *testing.T) {
	f := math.Sqrt2 / 2
	w := newDefaultWorld()
	s := NewPlane()
	s.Material.Reflective = 0.5
	s.Transform = alg.Translation(0, -1, 0)
	w.AddObject(s)
	r := NewRay(alg.NewPoint(0, 0, -3), alg.NewVector3(0, -f, f))
	i := NewIntersection(math.Sqrt2, s)
	c := prepareComps(i, r, NewIntersections(i))
	color := w.reflectedColor(c, w.MaxBounces)
	AssertColorEqual(t, color, canvas.NewColor(0.190321, 0.237913, 0.142748))
}

func TestTheShadeHitWithReflectiveMaterial(t *testing.T) {
	f := math.Sqrt2 / 2
	w := newDefaultWorld()
	s := NewPlane()
	s.Material.Reflective = 0.5
	s.Transform = alg.Translation(0, -1, 0)
	w.AddObject(s)
	r := NewRay(alg.NewPoint(0, 0, -3), alg.NewVector3(0, -f, f))
	i := NewIntersection(math.Sqrt2, s)
	c := prepareComps(i, r, NewIntersections(i))
	color := w.shade(c, w.MaxBounces)
	AssertColorEqual(t, color, canvas.NewColor(0.876756, 0.924339, 0.829173))
}

// TODO: Test for termination
// func TestAvoidInfiniteRecursionForReflection(t *testing.T) {
// 	w := NewWorld()
// 	l := NewPointLight(alg.NewPoint(0, 0, 0), canvas.White)
// 	p1 := NewPlane()
// 	p1.Material.Reflective = 1
// 	p1.Transform = alg.Translation(0, -1, 0)
// 	p2 := NewPlane()
// 	p2.Material.Reflective = 1
// 	p2.Transform = alg.Translation(0, 1, 0)
// 	w.AddLight(l)
// 	w.AddObject(p1)
// 	w.AddObject(p2)
// 	r := NewRay(alg.NewPoint(0, 0, 0), alg.NewVector3(0, 1, 0))
// 	color := w.colorAt(r, w.MaxBounces)
// 	AssertColorEqual(t, color, canvas.NewColor(0.876756, 0.924339, 0.829173))
// }

func TestReflectedColorAtMaximumDepth(t *testing.T) {
	f := math.Sqrt2 / 2
	w := newDefaultWorld()
	s := NewPlane()
	s.Material.Reflective = 0.5
	s.Transform = alg.Translation(0, -1, 0)
	w.AddObject(s)
	r := NewRay(alg.NewPoint(0, 0, -3), alg.NewVector3(0, -f, f))
	i := NewIntersection(math.Sqrt2, s)
	c := prepareComps(i, r, NewIntersections(i))
	color := w.reflectedColor(c, 0)
	AssertColorEqual(t, color, canvas.Black)
}

func TestTheRefractedColorWithAnOpaqueSurface(t *testing.T) {
	w := newDefaultWorld()
	s := w.Objects[0]
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	xs := NewIntersections(
		NewIntersection(4, s),
		NewIntersection(6, s),
	)
	c := prepareComps(xs[0], r, xs)
	color := w.refractedColor(c, w.MaxBounces)
	AssertColorEqual(t, color, canvas.Black)
}

func TestRefractedColorAtMaximumDepth(t *testing.T) {
	w := newDefaultWorld()
	s := w.Objects[0]
	s.GetMaterial().Transparency = 1.0
	s.GetMaterial().RefractiveIndex = 1.5
	r := NewRay(alg.NewPoint(0, 0, -5), alg.NewVector3(0, 0, 1))
	xs := NewIntersections(
		NewIntersection(4, s),
		NewIntersection(6, s),
	)
	c := prepareComps(xs[0], r, xs)
	color := w.refractedColor(c, 0)
	AssertColorEqual(t, color, canvas.Black)
}

func TestRefractedColorUnderTotalInternalReflection(t *testing.T) {
	f := math.Sqrt2 / 2
	w := newDefaultWorld()
	s := w.Objects[0]
	s.GetMaterial().Transparency = 1.0
	s.GetMaterial().RefractiveIndex = 1.5
	r := NewRay(alg.NewPoint(0, 0, f), alg.NewVector3(0, 1, 0))
	xs := NewIntersections(
		NewIntersection(-f, s),
		NewIntersection(f, s),
	)
	// NOTE: We are inside the sphere so we need
	// to look at the second intersection xs[1].
	c := prepareComps(xs[1], r, xs)
	color := w.refractedColor(c, 5)
	AssertColorEqual(t, color, canvas.Black)
}

func TestRefractedColorWithARefractedRay(t *testing.T) {
	w := newDefaultWorld()
	a := w.Objects[0]
	a.GetMaterial().Ambient = 1.0
	a.GetMaterial().Pattern = newTestPattern()
	b := w.Objects[1]
	b.GetMaterial().Transparency = 1.0
	b.GetMaterial().RefractiveIndex = 1.5
	r := NewRay(alg.NewPoint(0, 0, 0.1), alg.NewVector3(0, 1, 0))
	xs := NewIntersections(
		NewIntersection(-0.9899, a),
		NewIntersection(-0.4899, b),
		NewIntersection(0.4899, b),
		NewIntersection(0.9899, a),
	)
	// NOTE: We are inside the sphere so we need
	// to look at the second intersection xs[1].
	c := prepareComps(xs[2], r, xs)
	color := w.refractedColor(c, 5)
	AssertColorEqual(t, color, canvas.NewColor(0, 0.998885, 0.04725))
}

func TestShadeHitWithATransparentMaterial(t *testing.T) {
	f := math.Sqrt2 / 2
	w := newDefaultWorld()
	floor := NewPlane()
	floor.Transform = alg.Translation(0, -1, 0)
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5
	w.AddObject(floor)
	ball := NewSphere()
	ball.Transform = alg.Translation(0, -3.5, -0.5)
	ball.Material.Color = canvas.NewColor(1, 0, 0)
	ball.Material.Ambient = 0.5
	w.AddObject(ball)
	r := NewRay(alg.NewPoint(0, 0, -3), alg.NewVector3(0, -f, f))
	xs := NewIntersections(
		NewIntersection(math.Sqrt2, floor),
	)
	c := prepareComps(xs[0], r, xs)
	color := w.shade(c, w.MaxBounces)
	AssertColorEqual(t, color, canvas.NewColor(0.93642, 0.686425, 0.686425))
}

func newDefaultWorld() *World {
	w := NewWorld()
	lp := alg.NewPoint(-10, 10, -10)
	l := NewPointLight(lp, canvas.White)
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

func newDefaultWorldWithoutLight() *World {
	w := NewWorld()
	s1 := NewSphere()
	s1.Material.Color = canvas.NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2 := NewSphere()
	s2.Transform = alg.Scaling(0.5, 0.5, 0.5)
	w.AddObject(s1)
	w.AddObject(s2)
	return w
}
