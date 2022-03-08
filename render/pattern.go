package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

// https://go.dev/doc/effective_go#embedding

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
	x := math.Floor(patternPoint[0])
	if math.Mod(x, 2) == 0 {
		return p.a.ColorAt(object, patternPoint)
	}
	return p.b.ColorAt(object, patternPoint)
}

type Checkers2DPattern struct {
	a         Pattern
	b         Pattern
	Transform alg.Matrix
}

func NewCheckers2DPattern(a Pattern, b Pattern) *Checkers2DPattern {
	return &Checkers2DPattern{
		a:         a,
		b:         b,
		Transform: alg.Id4,
	}
}

func (p *Checkers2DPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	x := math.Floor(patternPoint[0])
	y := math.Floor(patternPoint[1])
	if math.Mod(x+y, 2) == 0 {
		return p.a.ColorAt(object, patternPoint)
	}
	return p.b.ColorAt(object, patternPoint)
}

type Checkers3DPattern struct {
	a         Pattern
	b         Pattern
	Transform alg.Matrix
}

func NewCheckers3DPattern(a Pattern, b Pattern) *Checkers3DPattern {
	return &Checkers3DPattern{
		a:         a,
		b:         b,
		Transform: alg.Id4,
	}
}

func (p *Checkers3DPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	x := math.Floor(patternPoint[0])
	y := math.Floor(patternPoint[1])
	z := math.Floor(patternPoint[2])
	if math.Mod(x+y+z, 2) == 0 {
		return p.a.ColorAt(object, patternPoint)
	}
	return p.b.ColorAt(object, patternPoint)
}

type GradientPattern struct {
	a         Pattern
	b         Pattern
	Transform alg.Matrix
}

func NewGradientPattern(a Pattern, b Pattern) *GradientPattern {
	return &GradientPattern{
		a:         a,
		b:         b,
		Transform: alg.Id4,
	}
}

func (p *GradientPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	fraction := patternPoint[0] - math.Floor(patternPoint[0])
	colora := p.a.ColorAt(object, patternPoint).Scale(fraction)
	colorb := p.b.ColorAt(object, patternPoint).Scale(1 - fraction)
	return colora.Add(colorb)
}

type BlendedPattern struct {
	patterns  []Pattern
	Transform alg.Matrix
}

func NewBlendedPattern(patterns ...Pattern) *BlendedPattern {
	return &BlendedPattern{
		patterns:  patterns,
		Transform: alg.Id4,
	}
}

func (p *BlendedPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	color := canvas.Black
	for _, pattern := range p.patterns {
		color = color.Add(pattern.ColorAt(object, patternPoint))
	}
	avg := 1.0 / float64(len(p.patterns))
	return color.Scale(avg)
}

type RingPattern struct {
	a         Pattern
	b         Pattern
	Transform alg.Matrix
}

func NewRingPattern(a Pattern, b Pattern) *RingPattern {
	return &RingPattern{
		a:         a,
		b:         b,
		Transform: alg.Id4,
	}
}

func (p *RingPattern) ColorAt(object Object, point alg.Vector) canvas.Color {
	objectPoint := alg.Inverse(object.GetTransform()).MultVec(point)
	patternPoint := alg.Inverse(p.Transform).MultVec(objectPoint)
	x := patternPoint[0]
	y := patternPoint[1]
	d := math.Sqrt(x*x + y*y)
	if math.Mod(math.Floor(d), 2) == 0 {
		return p.a.ColorAt(object, patternPoint)
	}
	return p.b.ColorAt(object, patternPoint)
}
