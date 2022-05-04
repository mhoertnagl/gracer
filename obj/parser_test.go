package obj

import (
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/render"
)

func TestIgnoreUnrecognizedLines(t *testing.T) {
	p := NewParser()
	p.ParseString(`
		There was a young lady named Bright
		who traveled much faster than light.
		She set out one day
		in a relative way,
		and came back the previous night.
	`)
}

func TestVertex(t *testing.T) {
	p := NewParser()
	p.ParseString(`
		v -1 1 0
		v -1.0000 0.5000 0.000
		v 1 0 0
		v 1 1 0
	`)
	AssertIntEqual(t, len(p.Vertices), 4)
	AssertVectorEqual(t, p.getVertex(1), alg.NewPoint(-1, 1, 0))
	AssertVectorEqual(t, p.getVertex(2), alg.NewPoint(-1, 0.5, 0))
	AssertVectorEqual(t, p.getVertex(3), alg.NewPoint(1, 0, 0))
	AssertVectorEqual(t, p.getVertex(4), alg.NewPoint(1, 1, 0))
}

func TestTriangleFaces(t *testing.T) {
	p := NewParser()
	p.ParseString(`
		v -1 1 0
		v -1.0000 0.5000 0.000
		v 1 0 0
		v 1 1 0

		f 1 2 3
		f 1 3 4
	`)
	AssertIntEqual(t, len(p.Root.Kids), 2)
	AssertVectorEqual(t, p.getTriangle(0).P1, p.getVertex(1))
	AssertVectorEqual(t, p.getTriangle(0).P2, p.getVertex(2))
	AssertVectorEqual(t, p.getTriangle(0).P3, p.getVertex(3))
	AssertVectorEqual(t, p.getTriangle(1).P1, p.getVertex(1))
	AssertVectorEqual(t, p.getTriangle(1).P2, p.getVertex(3))
	AssertVectorEqual(t, p.getTriangle(1).P3, p.getVertex(4))
}

func TestTriangulatingPolygons(t *testing.T) {
	p := NewParser()
	p.ParseString(`
		v -1 1 0
		v -1 0 0
		v 1 0 0
		v 1 1 0
		v 0 2 0

		f 1 2 3 4 5
	`)
	AssertIntEqual(t, len(p.Root.Kids), 3)
	AssertVectorEqual(t, p.getTriangle(0).P1, p.getVertex(1))
	AssertVectorEqual(t, p.getTriangle(0).P2, p.getVertex(2))
	AssertVectorEqual(t, p.getTriangle(0).P3, p.getVertex(3))
	AssertVectorEqual(t, p.getTriangle(1).P1, p.getVertex(1))
	AssertVectorEqual(t, p.getTriangle(1).P2, p.getVertex(3))
	AssertVectorEqual(t, p.getTriangle(1).P3, p.getVertex(4))
	AssertVectorEqual(t, p.getTriangle(2).P1, p.getVertex(1))
	AssertVectorEqual(t, p.getTriangle(2).P2, p.getVertex(4))
	AssertVectorEqual(t, p.getTriangle(2).P3, p.getVertex(5))
}

func TestTrianglesInGroup(t *testing.T) {
	p := NewParser()
	p.ParseString(`
		v -1 1 0
		v -1 0 0
		v 1 0 0
		v 1 1 0

		g FirstGroup
		f 1 2 3
		g SecondGroup
		f 1 3 4
	`)

	AssertIntEqual(t, len(p.Root.Kids), 2)
	g1 := p.Root.Kids[0].(*render.Group)
	g2 := p.Root.Kids[1].(*render.Group)

	AssertIntEqual(t, len(g1.Kids), 1)
	t1 := g1.Kids[0].(*render.Triangle)
	AssertVectorEqual(t, t1.P1, p.getVertex(1))
	AssertVectorEqual(t, t1.P2, p.getVertex(2))
	AssertVectorEqual(t, t1.P3, p.getVertex(3))

	AssertIntEqual(t, len(g2.Kids), 1)
	t2 := g2.Kids[0].(*render.Triangle)
	AssertVectorEqual(t, t2.P1, p.getVertex(1))
	AssertVectorEqual(t, t2.P2, p.getVertex(3))
	AssertVectorEqual(t, t2.P3, p.getVertex(4))
}
