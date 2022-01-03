package alg

type Matrix []Vector

// TODO: special type for square matrices.
// type squareMatrix = []vector

var Id2 = NewIdMatrix(2)
var Id3 = NewIdMatrix(3)
var Id4 = NewIdMatrix(4)

func NewZeroMatrix(size int) Matrix {
	m := make(Matrix, size)
	for i := range m {
		m[i] = NewZeroVector(size)
	}
	return m
}

func NewIdMatrix(size int) Matrix {
	m := NewZeroMatrix(size)
	for i := 0; i < size; i++ {
		m[i][i] = 1.0
	}
	return m
}

func NewMatrix(rs ...Vector) Matrix {
	return rs
}

func Row(vs ...float64) Vector {
	return vs
}

func (a Matrix) MultMat(b Matrix) Matrix {
	m := NewZeroMatrix(len(a))
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a); c++ {
			for l := 0; l < len(a); l++ {
				m[r][c] += a[r][l] * b[l][c]
			}
		}
	}
	return m
}

func (a Matrix) MultVec(b Vector) Vector {
	v := NewZeroVector(len(a))
	for r := 0; r < len(a); r++ {
		for l := 0; l < len(a); l++ {
			v[r] += a[r][l] * b[l]
		}
	}
	return v
}

func (a Matrix) Transpose() Matrix {
	m := NewZeroMatrix(len(a))
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a); c++ {
			m[c][r] = a[r][c]
		}
	}
	return m
}
