package alg

import (
	"math"

	"github.com/mhoertnagl/gracer/utils"
)

// Max returns the component-wise maximum vector of vectors a and b.
func Max(a, b Vector) Vector {
	max := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		max[i] = math.Max(a[i], b[i])
	}
	return max
}

// Min returns the component-wise minimum vector of vectors a and b.
func Min(a, b Vector) Vector {
	min := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		min[i] = math.Min(a[i], b[i])
	}
	return min
}

// Max3 returns the component-wise maximum vector of vectors a, b and c.
func Max3(a, b, c Vector) Vector {
	max := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		max[i] = utils.Max3(a[i], b[i], c[i])
	}
	return max
}

// Min3 returns the component-wise minimum vector of vectors a, b and c.
func Min3(a, b, c Vector) Vector {
	min := NewZeroVector(len(a))
	for i := 0; i < len(a); i++ {
		min[i] = utils.Min3(a[i], b[i], c[i])
	}
	return min
}

// MaxN returns the component-wise maximum vector of vectors as.
func MaxN(as ...Vector) Vector {
	if len(as) == 0 {
		return NewZeroVector(4)
	}
	size := len(as[0])
	max := NewValueVector(size, -math.MaxFloat64)
	for _, a := range as {
		for i := 0; i < size; i++ {
			max[i] = math.Max(max[i], a[i])
		}
	}
	return max
}

// MinN returns the component-wise minimum vector of vectors as
func MinN(as ...Vector) Vector {
	if len(as) == 0 {
		return NewZeroVector(4)
	}
	size := len(as[0])
	min := NewValueVector(size, math.MaxFloat64)
	for _, a := range as {
		for i := 0; i < size; i++ {
			min[i] = math.Max(min[i], a[i])
		}
	}
	return min
}
