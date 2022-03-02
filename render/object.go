package render

import "github.com/mhoertnagl/gracer/alg"

type Object interface {
	Intersect(r *Ray) Intersections
	NormalAt(p alg.Vector) alg.Vector
	GetMaterial() *Material
}
