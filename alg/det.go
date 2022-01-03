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
	{1, 2, 3, 4}, // even
	{1, 2, 4, 3}, // odd
	{1, 3, 4, 2}, // even
	{2, 3, 4, 1}, // odd
	{2, 3, 1, 4}, // even
	{2, 4, 1, 3}, // odd
	{3, 4, 1, 2}, // even
	{3, 4, 2, 1}, // odd
	{3, 1, 2, 4}, // even
	{4, 1, 2, 3}, // odd
	{4, 1, 3, 2}, // even
	{4, 2, 3, 1}, // odd
	// swap 2, 3
	{4, 3, 2, 1}, // even
	{4, 3, 1, 2}, // odd
	{4, 2, 1, 3}, // even
	{3, 2, 1, 4}, // odd
	{3, 2, 4, 1}, // even
	{3, 1, 4, 2}, // odd
	{2, 1, 4, 3}, // even
	{2, 1, 3, 4}, // odd
	{2, 4, 3, 1}, // even
	{1, 4, 3, 2}, // odd
	{1, 4, 2, 3}, // even
	{1, 3, 2, 4}, // odd
}

func det(tab ptable, m Matrix) float64 {
	sz := len(m)
	sign := 1.0
	result := 0.0
	for _, p := range tab {
		factor := m[0][p[0]]
		for i := 1; i < sz; i++ {
			factor *= m[i][p[i]]
		}
		result += sign * factor
		sign *= -1.0
	}
	return result
}

func det2(m Matrix) float64 {
	return m[0][0]*m[1][1] - m[1][0]*m[0][1]
}

func det3(m Matrix) float64 {
	return det(ptable3, m)
}

func det4(m Matrix) float64 {
	return det(ptable4, m)
}

func (m Matrix) Det() float64 {
	switch len(m) {
	case 0:
		return 0
	case 1:
		return m[0][0]
	case 2:
		return det2(m)
	case 3:
		return det3(m)
	case 4:
		return det4(m)
	}
	panic("Cannot compute determinant of matrix larger than 4")
}
