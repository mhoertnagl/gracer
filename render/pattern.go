package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

type Pattern interface {
	ColorAt(object Object, point alg.Vector) canvas.Color
}

type SolidPattern struct {
	color canvas.Color
}

// TODO: Pass Material
func NewSolidPattern(color canvas.Color) *SolidPattern {
	return &SolidPattern{color: color}
}

func (p *SolidPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	return p.color
}

type StripePattern struct {
	a         Pattern
	b         Pattern
	Transform alg.Matrix
}

func NewStripePattern(a Pattern, b Pattern) *StripePattern {
	return &StripePattern{
		a:         a,
		b:         b,
		Transform: alg.Id4,
	}
}

func (p *StripePattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	if math.Mod(math.Floor(patternPoint[0]), 2) == 0 {
		return p.a.ColorAt(object, patternPoint)
	}
	return p.b.ColorAt(object, patternPoint)
}
