package render

import (
	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

type PointLight struct {
	Position  alg.Vector
	Intensity canvas.Color
}

func NewPointLight(pos alg.Vector, intensity canvas.Color) *PointLight {
	return &PointLight{Position: pos, Intensity: intensity}
}
