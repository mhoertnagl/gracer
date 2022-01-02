package alg

type vector = []float64

func newZeroVector(size int) vector {
	return make(vector, size)
}

func newVector(vs ...float64) vector {
	return vs
}
