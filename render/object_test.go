package render

import (
	"math"
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestConvertingAPointFromWorldToObjectSpace(t *testing.T) {
	g1 := NewGroup()
	g1.Transform = alg.RotationY(math.Pi / 2)
	g2 := NewGroup()
	g2.Transform = alg.Scaling(2, 2, 2)
	g1.AddKid(g2)
	s := NewSphere()
	s.Transform = alg.Translation(5, 0, 0)
	g2.AddKid(s)
	p := worldToObject(s, alg.NewPoint(-2, 0, -10))
	AssertVectorEqual(t, p, alg.NewPoint(0, 0, -1))
}

func TestConvertingANormalFromObjectToWorldSpace(t *testing.T) {
	f := math.Sqrt(3) / 3
	g1 := NewGroup()
	g1.Transform = alg.RotationY(math.Pi / 2)
	g2 := NewGroup()
	g2.Transform = alg.Scaling(1, 2, 3)
	g1.AddKid(g2)
	s := NewSphere()
	s.Transform = alg.Translation(5, 0, 0)
	g2.AddKid(s)
	n := normalToWorld(s, alg.NewVector3(f, f, f))
	AssertVectorEqual(t, n, alg.NewVector3(0.285714, 0.428571, -0.857143))
}

func TestFindingTheNormalOnAChildObject(t *testing.T) {
	g1 := NewGroup()
	g1.Transform = alg.RotationY(math.Pi / 2)
	g2 := NewGroup()
	g2.Transform = alg.Scaling(1, 2, 3)
	g1.AddKid(g2)
	s := NewSphere()
	s.Transform = alg.Translation(5, 0, 0)
	g2.AddKid(s)
	n := s.NormalAt(alg.NewPoint(1.7321, 1.1547, -5.5774), nil)
	AssertVectorEqual(t, n, alg.NewVector3(0.285704, 0.428543, -0.857161))
}
