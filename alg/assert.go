package alg

import (
	"math"
	"testing"
)

func AssertFloatEqual(t *testing.T, a, e float64) {
	t.Helper()
	if !floatEqual(a, e) {
		t.Errorf("Float was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func AssertVectorEqual(t *testing.T, a, e Vector) {
	t.Helper()
	if !vectorEqual(a, e) {
		t.Errorf("Vector was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func AssertMatrixEqual(t *testing.T, a, e Matrix) {
	t.Helper()
	if len(a) != len(e) {
		t.Errorf("Matrix size was incorrect, \n got: %v \n want: %v", len(a), len(e))
	} else if !matrixEqual(a, e) {
		t.Errorf("Matrix was incorrect, \n got: %v \n want: %v", a, e)
	}
}

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}

func vectorEqual(a, b Vector) bool {
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

func matrixEqual(a, b Matrix) bool {
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
