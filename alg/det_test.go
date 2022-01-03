package alg

import "testing"

func TestDeterminantOf2x2Matrix(t *testing.T) {
	a := NewMatrix(
		Row(1, 5),
		Row(-3, 2),
	)
	AssertFloatEqual(t, a.Det(), 17)
}
