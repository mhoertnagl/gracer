package render

import (
	"math"
	"testing"

	"github.com/mhoertnagl/gracer/alg"
)

func TestPixelSizeForAHorizontalCanvas(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)
	AssertFloatEqual(t, c.pixelSize, 0.01)
}

func TestPixelSizeForAVerticalCanvas(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)
	AssertFloatEqual(t, c.pixelSize, 0.01)
}

func TestAViewTransformLookingInPositiveZDirection(t *testing.T) {
	from := alg.Origin
	to := alg.NewPoint(0, 0, 1)
	up := alg.NewVector3(0, 1, 0)
	vt := ViewTransform(from, to, up)
	AssertMatrixEqual(t, vt, alg.Scaling(-1, 1, -1))
}

func TestTheViewTransformMovesTheWorld(t *testing.T) {
	from := alg.NewPoint(0, 0, 8)
	to := alg.Origin
	up := alg.NewVector3(0, 1, 0)
	vt := ViewTransform(from, to, up)
	AssertMatrixEqual(t, vt, alg.Translation(0, 0, -8))
}

func TestAnWribtraryViewTransform(t *testing.T) {
	e := alg.NewMatrix(
		alg.Row(-0.50709, 0.50709, 0.67612, -2.36643),
		alg.Row(0.76772, 0.60609, 0.12122, -2.82843),
		alg.Row(-0.35857, 0.59761, -0.71714, 0),
		alg.Row(0, 0, 0, 1),
	)
	from := alg.NewPoint(1, 3, 2)
	to := alg.NewPoint(4, -2, 8)
	up := alg.NewVector3(1, 1, 0)
	vt := ViewTransform(from, to, up)
	AssertMatrixEqual(t, vt, e)
}

func TestConstructingARayThroughTheCenterOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	AssertVectorEqual(t, r.Origin, alg.Origin)
	AssertVectorEqual(t, r.Direction, alg.NewVector3(0, 0, -1))
}

func TestConstructingARayThroughTheCornerOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)
	AssertVectorEqual(t, r.Origin, alg.Origin)
	AssertVectorEqual(t, r.Direction, alg.NewVector3(0.66519, 0.33259, -0.66851))
}

func TestConstructingARayWhenTheCameraIsTransformed(t *testing.T) {
	f := math.Sqrt2 / 2
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = alg.RotationY(math.Pi / 4).MultMat(alg.Translation(0, -2, 5))
	r := c.RayForPixel(100, 50)
	AssertVectorEqual(t, r.Origin, alg.NewPoint(0, 2, -5))
	AssertVectorEqual(t, r.Direction, alg.NewVector3(f, 0, -f))
}
