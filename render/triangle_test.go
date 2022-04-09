package render

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestConstructingATriangle(t *testing.T) {
	p1 := alg.NewPoint(0, 1, 0)
	p2 := alg.NewPoint(-1, 0, 0)
	p3 := alg.NewPoint(1, 0, 0)
	tr := NewTriangle(p1, p2, p3)
	AssertVectorEqual(t, tr.P1, p1)
	AssertVectorEqual(t, tr.P2, p2)
	AssertVectorEqual(t, tr.P3, p3)
	AssertVectorEqual(t, tr.E1, alg.NewVector3(-1, -1, 0))
	AssertVectorEqual(t, tr.E2, alg.NewVector3(1, -1, 0))
	AssertVectorEqual(t, tr.Normal, alg.NewVector3(0, 0, -1))
}
