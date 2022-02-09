package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
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

func (m *Material) Lighting(l *PointLight, p alg.Vector, ev alg.Vector, nv alg.Vector) canvas.Color {
	ec := m.Color.Mul(l.Intensity)
	lv := l.Position.Sub(p).Norm()
	amb := ec.Scale(m.Ambient)
	dif := canvas.Black()
	spe := canvas.Black()
	ldn := lv.Dot(nv)
	if ldn > 0 {
		dif = ec.Scale(m.Diffuse * ldn)
		rv := lv.Neg().Reflect(nv)
		rde := rv.Dot(ev)
		if rde > 0 {
			f := math.Pow(rde, m.Shininess)
			spe = l.Intensity.Scale(f * m.Specular)
		}
	}
	return amb.Add(dif).Add(spe)
}
