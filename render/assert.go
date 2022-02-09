package render

import (
	"math"
	"testing"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
)

func AssertIntEqual(t *testing.T, a, e int) {
	t.Helper()
	if a != e {
		t.Errorf("Int was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func AssertFloatEqual(t *testing.T, a, e float64) {
	t.Helper()
	if !floatEqual(a, e) {
		t.Errorf("Float was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func AssertVectorEqual(t *testing.T, a, e alg.Vector) {
	t.Helper()
	if !vectorEqual(a, e) {
		t.Errorf("Vector was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func AssertMatrixEqual(t *testing.T, a, e alg.Matrix) {
	t.Helper()
	if len(a) != len(e) {
		t.Errorf("Matrix size was incorrect, \n got: %v \n want: %v", len(a), len(e))
	} else if !matrixEqual(a, e) {
		t.Errorf("Matrix was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func AssertColorEqual(t *testing.T, a, e canvas.Color) {
	t.Helper()
	if !colorEqual(a, e) {
		t.Errorf("Color was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-5
}

func vectorEqual(a, b alg.Vector) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !floatEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func matrixEqual(a, b alg.Matrix) bool {
	if len(a) != len(b) {
		return false
	}
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a); c++ {
			va := a[r][c]
			vb := b[r][c]
			if !floatEqual(va, vb) {
				return false
			}
		}
	}
	return true
}

func colorEqual(a, b canvas.Color) bool {
	return floatEqual(a[0], b[0]) &&
		floatEqual(a[1], b[1]) &&
		floatEqual(a[2], b[2])
}
