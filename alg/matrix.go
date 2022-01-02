package alg

type matrix []vector

// TODO: special type for square matrices.
// type squareMatrix = []vector

var id2 = newIdMatrix(2)
var id3 = newIdMatrix(3)
var id4 = newIdMatrix(4)

func newZeroMatrix(size int) matrix {
	m := make(matrix, size)
	for i := range m {
		m[i] = newZeroVector(size)
	}
	return m
}

func newIdMatrix(size int) matrix {
	m := newZeroMatrix(size)
	for i := 0; i < size; i++ {
		m[i][i] = 1.0
	}
	return m
}

func newMatrix(rs ...vector) matrix {
	return rs
}

func r(vs ...float64) vector {
	return vs
}
