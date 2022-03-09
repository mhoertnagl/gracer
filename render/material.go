package render

import (
	"github.com/mhoertnagl/gracer/canvas"
)

type Material struct {
	Color          canvas.Color
	Pattern        Pattern
	Ambient        float64
	Diffuse        float64
	Specular       float64
	Shininess      float64
	Reflective     float64
	Transparency   float64
	RefrctiveIndex float64
}

func NewMaterial() *Material {
	return &Material{
		Color:          canvas.White,
		Pattern:        nil,
		Ambient:        0.1,
		Diffuse:        0.9,
		Specular:       0.9,
		Shininess:      200.0,
		Reflective:     0.0,
		Transparency:   0.0,
		RefrctiveIndex: 1.0,
	}
}
