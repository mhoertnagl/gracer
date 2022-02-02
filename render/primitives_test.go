package render

import (
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
