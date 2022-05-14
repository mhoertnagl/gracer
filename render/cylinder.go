package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Cylinder struct {
	Transform    alg.Matrix
	Material     *Material
	Parent       Object
	invTransform alg.Matrix
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		Transform:    alg.Id4,
		Material:     NewMaterial(),
		Parent:       nil,
		invTransform: nil,
	}
}

func (s *Cylinder) Intersect(r *Ray) Intersections {
	r2 := r.Transform(s.GetInverseTransform())
	xd := r2.Direction[0]
	zd := r2.Direction[2]
	a := xd*xd + zd*zd
	if math.Abs(a) < EPSILON {
		return NewIntersections()
	}
	xo := r2.Origin[0]
	zo := r2.Origin[2]
	b := xo*xd + zo*zd
	c := xo*xo + zo*zo - 1
	d := b*b - a*c
	if d < 0 {
		return NewIntersections()
	}
	ds := math.Sqrt(d)
	t0 := (-b - ds) / a
	t1 := (-b + ds) / a
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	yo := r2.Origin[1]
	yd := r2.Direction[1]
	xs := NewIntersections()
	y0 := yo + t0*yd
	if 0 <= y0 && y0 <= 1 {
		xs = append(xs, NewIntersection(t0, s))
	}
	y1 := yo + t1*yd
	if 0 <= y1 && y1 <= 1 {
		xs = append(xs, NewIntersection(t1, s))
	}
	if math.Abs(yd) > EPSILON {
		t2 := -yo / yd
		if checkCaps(r2, t2) {
			xs = append(xs, NewIntersection(t2, s))
		}
		t3 := (1 - yo) / yd
		if checkCaps(r2, t3) {
			xs = append(xs, NewIntersection(t3, s))
		}
	}
	return xs
}

func checkCaps(r *Ray, t float64) bool {
	x := r.Origin[0] + t*r.Direction[0]
	z := r.Origin[2] + t*r.Direction[2]
	return x*x+z*z <= 1
}

func (c *Cylinder) NormalAt(p alg.Vector, hit *Intersection) alg.Vector {
	op := worldToObject(c, p)
	on := c.localNormalAt(op)
	return normalToWorld(c, on)
}

func (c *Cylinder) localNormalAt(p alg.Vector) alg.Vector {
	dy := p[0]*p[0] + p[2]*p[2]
	if dy < 1 && p[1] >= 1-EPSILON {
		return alg.NewVector3(0, 1, 0)
	}
	if dy < 1 && p[1] <= EPSILON {
		return alg.NewVector3(0, -1, 0)
	}
	return alg.NewVector3(p[0], 0, p[2])
}

func (c *Cylinder) GetMaterial() *Material {
	return c.Material
}

func (c *Cylinder) GetTransform() alg.Matrix {
	return c.Transform
}

func (c *Cylinder) SetParent(obj Object) {
	c.Parent = obj
}

func (c *Cylinder) GetParent() Object {
	return c.Parent
}

func (c *Cylinder) GetInverseTransform() alg.Matrix {
	if c.invTransform == nil {
		c.invTransform = alg.Inverse(c.Transform)
	}
	return c.invTransform
}

func (c *Cylinder) GetBounds() *Bounds {
	return NewBounds(alg.NewPoint(-1, -1, -1), alg.NewPoint(1, 1, 1))
}

func (c *Cylinder) Includes(obj Object) bool {
	return c == obj
}
