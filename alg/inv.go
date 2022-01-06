package alg

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
