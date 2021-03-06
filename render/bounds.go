package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/utils"
)

type Bounds struct {
	Min alg.Vector
	Max alg.Vector
}

func NewBounds(min alg.Vector, max alg.Vector) *Bounds {
	return &Bounds{Min: min, Max: max}
}

func NewBoundsFrom(vs ...alg.Vector) *Bounds {
	min := alg.MinN(vs...)
	max := alg.MaxN(vs...)
	return NewBounds(min, max)
}

func (b *Bounds) Intersect(r *Ray) bool {
	xtmin, xtmax := checkBoundsAxis(b.Min[0], b.Max[0], r.Origin[0], r.Direction[0])
	ytmin, ytmax := checkBoundsAxis(b.Min[1], b.Max[1], r.Origin[1], r.Direction[1])
	ztmin, ztmax := checkBoundsAxis(b.Min[2], b.Max[2], r.Origin[2], r.Direction[2])
	tmin := utils.Max3(xtmin, ytmin, ztmin)
	tmax := utils.Min3(xtmax, ytmax, ztmax)
	return tmax >= tmin
}

// TODO: More general than checkBounds. Use this in the cube case.
func checkBoundsAxis(min, max, origin, direction float64) (float64, float64) {
	tmin_numerator := min - origin
	tmax_numerator := max - origin
	tmin := 0.0
	tmax := 0.0
	if math.Abs(direction) >= EPSILON {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		// TODO: inft cube.go -> utils/math.go
		tmin = inft(tmin_numerator)
		tmax = inft(tmax_numerator)
	}
	if tmin > tmax {
		return tmax, tmin
	}
	return tmin, tmax
}

func (b *Bounds) Edges() []alg.Vector {
	return []alg.Vector{
		alg.NewPoint(b.Min[0], b.Min[1], b.Min[2]),
		alg.NewPoint(b.Min[0], b.Min[1], b.Max[2]),
		alg.NewPoint(b.Min[0], b.Max[1], b.Min[2]),
		alg.NewPoint(b.Min[0], b.Max[1], b.Max[2]),
		alg.NewPoint(b.Max[0], b.Min[1], b.Min[2]),
		alg.NewPoint(b.Max[0], b.Min[1], b.Max[2]),
		alg.NewPoint(b.Max[0], b.Max[1], b.Min[2]),
		alg.NewPoint(b.Max[0], b.Max[1], b.Max[2]),
	}
}
