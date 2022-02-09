package render

import (
	"math"
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestASpheresDefaultTransformation(t *testing.T) {
	s := NewSphere()
	AssertMatrixEqual(t, s.Transform, alg.Id4)
}

func TestChangingASpheresDefaultTransformation(t *testing.T) {
	s := NewSphere()
	m := alg.Translation(2, 3, 4)
	s.Transform = m
	AssertMatrixEqual(t, s.Transform, m)
}

func TestTheNormalOnASphereAtAPointOnTheXAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(alg.NewPoint(1, 0, 0))
	AssertVectorEqual(t, n, alg.NewVector3(1, 0, 0))
}

func TestTheNormalOnASphereAtAPointOnTheYAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(alg.NewPoint(0, 1, 0))
	AssertVectorEqual(t, n, alg.NewVector3(0, 1, 0))
}

func TestTheNormalOnASphereAtAPointOnTheZAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(alg.NewPoint(0, 0, 1))
	AssertVectorEqual(t, n, alg.NewVector3(0, 0, 1))
}

func TestTheNormalOnASphereAtANonaxialPoint(t *testing.T) {
	f := math.Sqrt(3) / 3
	s := NewSphere()
	n := s.NormalAt(alg.NewPoint(f, f, f))
	AssertVectorEqual(t, n, alg.NewVector3(f, f, f))
	AssertVectorEqual(t, n, n.Norm())
}

func TestComputingTheNormalOnATranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.Transform = alg.Translation(0, 1, 0)
	n := s.NormalAt(alg.NewPoint(0, 1.70711, -0.70711))
	AssertVectorEqual(t, n, alg.NewVector3(0, 0.70711, -0.70711))
}

func TestComputingTheNormalOnATransformedSphere(t *testing.T) {
	f := math.Sqrt(2) / 2
	s := NewSphere()
	s.Transform = alg.Id4.Scale(1, 0.5, 1)
	n := s.NormalAt(alg.NewPoint(0, f, -f))
	AssertVectorEqual(t, n, alg.NewVector3(0, 0.97014, -0.24254))
}

func TestComputingTheNormalOnATransformedSphere2(t *testing.T) {
	f := math.Sqrt(2) / 2
	s := NewSphere()
	s.Transform = alg.Id4.Scale(1, 0.5, 1).RotateZ(math.Pi / 5)
	n := s.NormalAt(alg.NewPoint(0, f, -f))
	AssertVectorEqual(t, n, alg.NewVector3(-0.41499, 0.86207, -0.29089))
}
