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

func (a matrix) mulMat(b matrix) matrix {
	m := newZeroMatrix(len(a))
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a); c++ {
			for l := 0; l < len(a); l++ {
				m[r][c] += a[r][l] * b[l][c]
			}
		}
	}
	return m
}

func (a matrix) mulVec(b vector) vector {
	v := newZeroVector(len(a))
	for r := 0; r < len(a); r++ {
		for l := 0; l < len(a); l++ {
			v[r] += a[r][l] * b[l]
		}
	}
	return v
}

func (a matrix) transpose() matrix {
	m := newZeroMatrix(len(a))
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a); c++ {
			m[c][r] = a[r][c]
		}
	}
	return m
}
