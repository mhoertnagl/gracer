package alg

type ptable [][]int

var ptable3 = ptable{
	{0, 1, 2}, // even
	{0, 2, 1}, // odd
	{1, 2, 0}, // even
	{1, 0, 2}, // odd
	{2, 0, 1}, // even
	{2, 1, 0}, // odd
}

var ptable4 = ptable{
	{0, 1, 2, 3}, // even
	{0, 1, 3, 2}, // odd
	{0, 2, 3, 1}, // even
	{1, 2, 3, 0}, // odd
	{1, 2, 0, 3}, // even
	{1, 3, 0, 2}, // odd
	{2, 3, 0, 1}, // even
	{2, 3, 1, 0}, // odd
	{2, 0, 1, 3}, // even
	{3, 0, 1, 2}, // odd
	{3, 0, 2, 1}, // even
	{3, 1, 2, 0}, // odd
	// swap 2, 3
	{3, 2, 1, 0}, // even
	{3, 2, 0, 1}, // odd
	{3, 1, 0, 2}, // even
	{2, 1, 0, 3}, // odd
	{2, 1, 3, 0}, // even
	{2, 0, 3, 1}, // odd
	{1, 0, 3, 2}, // even
	{1, 0, 2, 3}, // odd
	{1, 3, 2, 0}, // even
	{0, 3, 2, 1}, // odd
	{0, 3, 1, 2}, // even
	{0, 2, 1, 3}, // odd
}

func Inverse(m Matrix) Matrix {
	n := NewZeroMatrix(len(m))
	d := Det(m)
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m); c++ {
			// Note implcit transpose.
			n[c][r] = Cofactor(m, r, c) / d
		}
	}
	return n
}

func Det(m Matrix) float64 {
	switch len(m) {
	// case 0:
	// 	return 0
	// case 1:
	// 	return m[0][0]
	case 2:
		return det2(m)
	case 3:
		return det3(m)
	case 4:
		return det4(m)
	// default:
	// 	return det5(m)
	default:
		panic("Not implemented.")
	}
}

func det2(m Matrix) float64 {
	return m[0][0]*m[1][1] - m[1][0]*m[0][1]
}

func det3(m Matrix) float64 {
	return tableDet(ptable3, m)
}

func det4(m Matrix) float64 {
	return tableDet(ptable4, m)
}

func tableDet(tab ptable, m Matrix) float64 {
	sz := len(m)
	sgn := 1.0
	res := 0.0
	for _, p := range tab {
		fac := m[0][p[0]]
		for i := 1; i < sz; i++ {
			fac *= m[i][p[i]]
		}
		res += sgn * fac
		sgn *= -1.0
	}
	return res
}

// func det5(m Matrix) float64 {
// 	det := 0.0
// 	for c := 0; c < len(m); c++ {
// 		det += m[0][c] * Cofactor(m, 0, c)
// 	}
// 	return det
// }

func Cofactor(m Matrix, r int, c int) float64 {
	minor := Minor(m, r, c)
	if (r+c)%2 == 0 {
		return minor
	}
	return -minor
}

func Minor(m Matrix, r int, c int) float64 {
	return Det(SubMatrix(m, r, c))
}

func SubMatrix(m Matrix, r int, c int) Matrix {
	n := NewZeroMatrix(len(m) - 1)
	// Copy sub matrix 0..r, 0..c
	for cr := 0; cr < r; cr++ {
		for cc := 0; cc < c; cc++ {
			n[cr][cc] = m[cr][cc]
		}
	}
	// Copy sub matrix 0..r, c..max
	for cr := 0; cr < r; cr++ {
		for cc := c; cc < len(m)-1; cc++ {
			n[cr][cc] = m[cr][cc+1]
		}
	}
	// Copy sub matrix r..max, 0..c
	for cr := r; cr < len(m)-1; cr++ {
		for cc := 0; cc < c; cc++ {
			n[cr][cc] = m[cr+1][cc]
		}
	}
	// Copy sub matrix r..max, c..max
	for cr := r; cr < len(m)-1; cr++ {
		for cc := c; cc < len(m)-1; cc++ {
			n[cr][cc] = m[cr+1][cc+1]
		}
	}
	return n
}
