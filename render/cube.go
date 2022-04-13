package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/utils"
)

type Cube struct {
	Transform    alg.Matrix
	Material     *Material
	Parent       Object
	invTransform alg.Matrix
}

func NewCube() *Cube {
	return &Cube{
		Transform:    alg.Id4,
		Material:     NewMaterial(),
		Parent:       nil,
		invTransform: nil,
	}
}

// TODO: Improve performance.
func (c *Cube) Intersect(r *Ray) Intersections {
	or := r.Transform(c.GetInverseTransform())
	xtmin, xtmax := checkAxis(or.Origin[0], or.Direction[0])
	ytmin, ytmax := checkAxis(or.Origin[1], or.Direction[1])
	ztmin, ztmax := checkAxis(or.Origin[2], or.Direction[2])
	tmin := utils.Max3(xtmin, ytmin, ztmin)
	tmax := utils.Min3(xtmax, ytmax, ztmax)
	if tmin > tmax {
		return NewIntersections()
	}
	return NewIntersections(
		NewIntersection(tmin, c),
		NewIntersection(tmax, c),
	)
}

func checkAxis(origin, direction float64) (float64, float64) {
	tmin_numerator := -1 - origin
	tmax_numerator := 1 - origin
	tmin := 0.0
	tmax := 0.0
	if math.Abs(direction) >= EPSILON {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		tmin = inft(tmin_numerator)
		tmax = inft(tmax_numerator)
	}
	if tmin > tmax {
		return tmax, tmin
	}
	return tmin, tmax
}

func inft(num float64) float64 {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return math.MaxFloat64
	}
	return -math.MaxFloat64
}

func (c *Cube) NormalAt(p alg.Vector) alg.Vector {
	op := worldToObject(c, p)
	on := c.localNormalAt(op)
	return normalToWorld(c, on)
}

func (c *Cube) localNormalAt(p alg.Vector) alg.Vector {
	xabs := math.Abs(p[0])
	yabs := math.Abs(p[1])
	zabs := math.Abs(p[2])
	maxc := utils.Max3(xabs, yabs, zabs)
	if maxc == xabs {
		return alg.NewVector3(p[0], 0, 0)
	}
	if maxc == yabs {
		return alg.NewVector3(0, p[1], 0)
	}
	return alg.NewVector3(0, 0, p[2])
}

func (c *Cube) GetMaterial() *Material {
	return c.Material
}

func (c *Cube) GetTransform() alg.Matrix {
	return c.Transform
}

func (c *Cube) SetParent(obj Object) {
	c.Parent = obj
}

func (c *Cube) GetParent() Object {
	return c.Parent
}

func (c *Cube) GetInverseTransform() alg.Matrix {
	if c.invTransform == nil {
		c.invTransform = alg.Inverse(c.Transform)
	}
	return c.invTransform
}

func (c *Cube) GetBounds() *Bounds {
	return NewBounds(alg.NewPoint(-1, -1, -1), alg.NewPoint(1, 1, 1))
}
