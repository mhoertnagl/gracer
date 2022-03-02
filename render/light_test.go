package render

import (
	"math"
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

func TestLightingWithTheEyeBetweenLightAndTheSurface(t *testing.T) {
	m := NewMaterial()
	p := alg.Origin
	ev := alg.NewVector3(0, 0, -1)
	nv := alg.NewVector3(0, 0, -1)
	lp := alg.NewPoint(0, 0, -10)
	li := canvas.NewColor(1, 1, 1)
	l := NewPointLight(lp, li)
	lc := l.Lighting(m, p, ev, nv, false)
	AssertColorEqual(t, lc, canvas.NewColor(1.9, 1.9, 1.9))
}

func TestLightingWithTheEyeBetweenLightAndTheSurfaceEyeOffset45Deg(t *testing.T) {
	f := math.Sqrt(2) / 2
	m := NewMaterial()
	p := alg.Origin
	ev := alg.NewVector3(0, f, -f)
	nv := alg.NewVector3(0, 0, -1)
	lp := alg.NewPoint(0, 0, -10)
	li := canvas.NewColor(1, 1, 1)
	l := NewPointLight(lp, li)
	lc := l.Lighting(m, p, ev, nv, false)
	AssertColorEqual(t, lc, canvas.NewColor(1, 1, 1))
}

func TestLightingWithTheEyeOppositeSurfaceLightOffset45Deg(t *testing.T) {
	m := NewMaterial()
	p := alg.Origin
	ev := alg.NewVector3(0, 0, -1)
	nv := alg.NewVector3(0, 0, -1)
	lp := alg.NewPoint(0, 10, -10)
	li := canvas.NewColor(1, 1, 1)
	l := NewPointLight(lp, li)
	lc := l.Lighting(m, p, ev, nv, false)
	AssertColorEqual(t, lc, canvas.NewColor(0.7364, 0.7364, 0.7364))
}

func TestLightingWithEyeInThePathOfTheReflectionVector(t *testing.T) {
	f := math.Sqrt(2) / 2
	m := NewMaterial()
	p := alg.Origin
	ev := alg.NewVector3(0, -f, -f)
	nv := alg.NewVector3(0, 0, -1)
	lp := alg.NewPoint(0, 10, -10)
	li := canvas.NewColor(1, 1, 1)
	l := NewPointLight(lp, li)
	lc := l.Lighting(m, p, ev, nv, false)
	AssertColorEqual(t, lc, canvas.NewColor(1.6364, 1.6364, 1.6364))
}

func TestLightingWithLightBehindSurface(t *testing.T) {
	m := NewMaterial()
	p := alg.Origin
	ev := alg.NewVector3(0, 0, -1)
	nv := alg.NewVector3(0, 0, -1)
	lp := alg.NewPoint(0, 0, 10)
	li := canvas.NewColor(1, 1, 1)
	l := NewPointLight(lp, li)
	lc := l.Lighting(m, p, ev, nv, false)
	AssertColorEqual(t, lc, canvas.NewColor(0.1, 0.1, 0.1))
}

func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	m := NewMaterial()
	p := alg.Origin
	ev := alg.NewVector3(0, 0, -1)
	nv := alg.NewVector3(0, 0, -1)
	lp := alg.NewPoint(0, 0, -10)
	li := canvas.NewColor(1, 1, 1)
	l := NewPointLight(lp, li)
	lc := l.Lighting(m, p, ev, nv, true)
	AssertColorEqual(t, lc, canvas.NewColor(0.1, 0.1, 0.1))
}

func TestThereIsNoShadowWhenNothingIsCollinearWhitAPointAndLight(t *testing.T) {
	w := newDefaultWorld()
	p := alg.NewPoint(0, 10, 0)
	AssertFalse(t, w.Lights[0].IsShadowed(w, p))
}

func TestThereIsAShadowWhenAnObjectIsBetweenThePointAndTheLight(t *testing.T) {
	w := newDefaultWorld()
	p := alg.NewPoint(10, -10, 10)
	AssertTrue(t, w.Lights[0].IsShadowed(w, p))
}

func TestThereIsNoShadowWhenAnObjectIsBehindTheLight(t *testing.T) {
	w := newDefaultWorld()
	p := alg.NewPoint(-20, 20, -20)
	AssertFalse(t, w.Lights[0].IsShadowed(w, p))
}

func TestThereIsNoShadowWhenAnObjectIsBehindThePoint(t *testing.T) {
	w := newDefaultWorld()
	p := alg.NewPoint(-2, 2, -2)
	AssertFalse(t, w.Lights[0].IsShadowed(w, p))
}
