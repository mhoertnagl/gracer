package render

import (
	"strings"

	"github.com/mhoertnagl/gracer/alg"
)

type Ray struct {
	origin    alg.Vector
	direction alg.Vector
}

func NewRay(origin, direction alg.Vector) Ray {
	return Ray{origin: origin, direction: direction}
}

func (r Ray) String() string {
	var b strings.Builder
	b.WriteString("Ray(")
	b.WriteString(r.origin.String())
	b.WriteString(", ")
	b.WriteString(r.direction.String())
	b.WriteString(")")
	return b.String()
}

func (r Ray) Position(t float64) alg.Vector {
	return r.origin.Add(r.direction.Mult(t))
}

func (r Ray) Transform(m alg.Matrix) Ray {
	return NewRay(m.MultVec(r.origin), m.MultVec(r.direction))
}
