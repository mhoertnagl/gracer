package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

func TestSolidPattern(t *testing.T) {
	c := canvas.NewColor(0.5, 0.7, 0.9)
	pattern := NewSolidPattern(c)
	s := NewSphere()

	as := []struct {
		p alg.Vector
		c canvas.Color
	}{
		{p: alg.NewPoint(0, 0, 0), c: c},
		{p: alg.NewPoint(0, 0, 1), c: c},
		{p: alg.NewPoint(1, 0, 0), c: c},
	}

	for _, a := range as {
		AssertColorEqual(t, pattern.ColorAt(s, a.p), a.c)
	}
}

func TestLightingWithSolidPattern(t *testing.T) {
	s := NewSphere()
	s.Material.Ambient = 1
	s.Material.Diffuse = 0
	s.Material.Specular = 0
	s.Material.Pattern = NewSolidPattern(canvas.NewColor(0.5, 0.7, 0.9))
	ev := alg.NewVector3(0, 0, -1)
	nv := alg.NewVector3(0, 0, -1)
	l := NewPointLight(alg.NewPoint(0, 0, -10), canvas.White)

	a := l.Lighting(s, alg.NewPoint(0, 0, 0), ev, nv, false)
	AssertColorEqual(t, a, canvas.NewColor(0.5, 0.7, 0.9))
}

func TestLAPatternWithAnObjectTransformation(t *testing.T) {
	s := NewSphere()
	s.Transform = alg.Scaling(2, 2, 2)
	pattern := newTestPattern()
	a := pattern.ColorAt(s, alg.NewPoint(2, 3, 4))
	AssertColorEqual(t, a, canvas.NewColor(1, 1.5, 2))
}

func TestLAPatternWithAPatternTransformation(t *testing.T) {
	s := NewSphere()
	pattern := newTestPattern()
	pattern.Transform = alg.Scaling(2, 2, 2)
	a := pattern.ColorAt(s, alg.NewPoint(2, 3, 4))
	AssertColorEqual(t, a, canvas.NewColor(1, 1.5, 2))
}

func TestLAPatternWithAnObjectAndAPatternTransformation(t *testing.T) {
	s := NewSphere()
	s.Transform = alg.Scaling(2, 2, 2)
	pattern := newTestPattern()
	pattern.Transform = alg.Translation(0.5, 1, 1.5)
	a := pattern.ColorAt(s, alg.NewPoint(2.5, 3, 3.5))
	AssertColorEqual(t, a, canvas.NewColor(0.75, 0.5, 0.25))
}

func TestStripePattern(t *testing.T) {
	solid1 := NewSolidPattern(canvas.White)
	solid2 := NewSolidPattern(canvas.Black)
	pattern := NewStripePattern(solid1, solid2)
	s := NewSphere()

	as := []struct {
		p alg.Vector
		c canvas.Color
	}{
		{p: alg.NewPoint(0, 0, 0), c: canvas.White},
		{p: alg.NewPoint(0, 0, 1), c: canvas.White},
		{p: alg.NewPoint(1, 0, 0), c: canvas.Black},
		{p: alg.NewPoint(2, 0, 0), c: canvas.White},
		{p: alg.NewPoint(2.9, 0, 0), c: canvas.White},
		{p: alg.NewPoint(3, 0, 0), c: canvas.Black},
		{p: alg.NewPoint(3.9, 0, 0), c: canvas.Black},
	}

	for _, a := range as {
		AssertColorEqual(t, pattern.ColorAt(s, a.p), a.c)
	}
}

type testPattern struct {
	Transform alg.Matrix
}

func newTestPattern() *testPattern {
	return &testPattern{Transform: alg.Id4}
}

func (p *testPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	return canvas.NewColor(patternPoint[0], patternPoint[1], patternPoint[2])
}
