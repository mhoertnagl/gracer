package alg

func (m Matrix) LuDecompose() (Matrix, Matrix) {
	sz := len(m)
	l := NewZeroMatrix(sz)
	u := NewZeroMatrix(sz)
	for i := 0; i < sz; i++ {
		for j := i; j < sz; j++ {
			d := 0.0
			e := 0.0
			for k := 0; k < i; k++ {
				d += l[j][k] * u[k][i]
				e += l[i][k] * u[k][j]
			}
			l[j][i] = m[j][i] - d
			u[i][j] = (m[i][j] - e) / l[i][i]
		}
	}
	return l, u
}
