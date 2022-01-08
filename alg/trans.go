package alg

func Translation(x, y, z float64) Matrix {
	return NewMatrix(
		Row(1, 0, 0, x),
		Row(0, 1, 0, y),
		Row(0, 0, 1, z),
		Row(0, 0, 0, 1),
	)
}

func Scaling(x, y, z float64) Matrix {
	return NewMatrix(
		Row(x, 0, 0, 0),
		Row(0, y, 0, 0),
		Row(0, 0, z, 0),
		Row(0, 0, 0, 1),
	)
}
