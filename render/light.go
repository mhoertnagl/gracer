package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

type Light interface {
	Lighting(m *Material, p, eye, normal alg.Vector, isShadowed bool) canvas.Color
	IsShadowed(w *World, p alg.Vector) bool
}

type PointLight struct {
	Position  alg.Vector
	Intensity canvas.Color
}

func NewPointLight(pos alg.Vector, intensity canvas.Color) *PointLight {
	return &PointLight{Position: pos, Intensity: intensity}
}

func (l *PointLight) Lighting(m *Material, p, eye, normal alg.Vector, isShadowed bool) canvas.Color {
	ec := m.Color.Mul(l.Intensity)
	amb := ec.Scale(m.Ambient)
	if isShadowed {
		return amb
	}
	lv := l.Position.Sub(p).Norm()
	dif := canvas.Black()
	spe := canvas.Black()
	ldn := lv.Dot(normal)
	if ldn > 0 {
		dif = ec.Scale(m.Diffuse * ldn)
		rv := lv.Neg().Reflect(normal)
		rde := rv.Dot(eye)
		if rde > 0 {
			f := math.Pow(rde, m.Shininess)
			spe = l.Intensity.Scale(f * m.Specular)
		}
	}
	return amb.Add(dif).Add(spe)
}

func (l *PointLight) IsShadowed(w *World, p alg.Vector) bool {
	v := l.Position.Sub(p)
	distance := v.Mag()
	direction := v.Norm()
	r := NewRay(p, direction)
	xxs := w.intersect(r)
	hit := xxs.Hit()
	return hit != nil && hit.Distance < distance
}
