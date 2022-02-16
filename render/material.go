package render

import (
	"github.com/mhoertnagl/gracer/canvas"
)

type Material struct {
	Color     canvas.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{
		Color:     canvas.NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
