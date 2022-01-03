package alg

import "testing"

func TestLuDecompose(t *testing.T) {
	el := NewMatrix(
		Row(3.0, 0, 0),
		Row(0.1, 7.003333, 0),
		Row(0.3, -0.19, 10.012042),
	)
	eu := NewMatrix(
		Row(1, -0.033333, -0.066667),
		Row(0, 1, -0.041885),
		Row(0, 0, 1),
	)
	x := NewMatrix(
		Row(3, -0.1, -0.2),
		Row(0.1, 7, -0.3),
		Row(0.3, -0.2, 10),
	)
	al, au := x.LuDecompose()
	AssertMatrixEqual(t, al, el)
	AssertMatrixEqual(t, au, eu)
	AssertMatrixEqual(t, al.MultMat(au), x)
}
