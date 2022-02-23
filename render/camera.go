package render

import (
	"math"

	"github.com/mhoertnagl/gracer/alg"
)

type Camera struct {
	hsize       int
	vsize       int
	fieldOfView float64
	halfWidth   float64
	halfHeight  float64
	pixelSize   float64
	Transform   alg.Matrix
}

func NewCamera(hsize, vsize int, fieldOfView float64) *Camera {
	c := &Camera{
		hsize:       hsize,
		vsize:       vsize,
		fieldOfView: fieldOfView,
		Transform:   alg.Id4,
	}
	halfView := math.Tan(c.fieldOfView / 2)
	aspect := float64(c.hsize) / float64(c.vsize)
	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfWidth = halfView * aspect
		c.halfHeight = halfView
	}
	c.pixelSize = 2 * c.halfWidth / float64(c.hsize)
	return c
}

func (c *Camera) RayForPixel(px, py int) *Ray {
	xoffset := (float64(px) + 0.5) * c.pixelSize
	yoffset := (float64(py) + 0.5) * c.pixelSize
	worldx := c.halfWidth - xoffset
	worldy := c.halfHeight - yoffset
	inv := alg.Inverse(c.Transform)
	pixel := inv.MultVec(alg.NewPoint(worldx, worldy, -1))
	origin := inv.MultVec(alg.Origin)
	direction := pixel.Sub(origin).Norm()
	return NewRay(origin, direction)
}

func ViewTransform(from, to, up alg.Vector) alg.Matrix {
	fwd := to.Sub(from).Norm()
	left := fwd.Cross(up.Norm())
	trueUp := left.Cross(fwd)
	orientation := alg.NewMatrix(
		alg.Row(left[0], left[1], left[2], 0),
		alg.Row(trueUp[0], trueUp[1], trueUp[2], 0),
		alg.Row(-fwd[0], -fwd[1], -fwd[2], 0),
		alg.Row(0, 0, 0, 1),
	)
	translation := alg.Translation(-from[0], -from[1], -from[2])
	return orientation.MultMat(translation)
}
