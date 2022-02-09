package render

import (
	"strings"

	"github.com/mhoertnagl/gracer/alg"
)

type Ray struct {
	Origin    alg.Vector
	Direction alg.Vector
}

func NewRay(origin, direction alg.Vector) *Ray {
	return &Ray{Origin: origin, Direction: direction}
}

func (r *Ray) String() string {
	var b strings.Builder
	b.WriteString("Ray(")
	b.WriteString(r.Origin.String())
	b.WriteString(", ")
	b.WriteString(r.Direction.String())
	b.WriteString(")")
	return b.String()
}

func (r *Ray) Position(t float64) alg.Vector {
	return r.Origin.Add(r.Direction.Mult(t))
}

func (r *Ray) Transform(m alg.Matrix) *Ray {
	return NewRay(m.MultVec(r.Origin), m.MultVec(r.Direction))
}
