package utils

import "math"

// Max3 returns the largest of x, y or z.
func Max3(x, y, z float64) float64 {
	return math.Max(x, math.Max(y, z))
}

// Min3 returns the smallest of x, y or z.
func Min3(x, y, z float64) float64 {
	return math.Min(x, math.Min(y, z))
}

// MaxN returns the largest of a list of numbers.
//
// Special cases are:
//  MaxN() = -Inf
//  MaxN(x) = x
func MaxN(xs ...float64) float64 {
	max := math.Inf(-1)
	for _, x := range xs {
		max = math.Max(max, x)
	}
	return max
}

// MinN returns the smallest of a list of numbers.
//
// Special cases are:
//  MinN() = +Inf
//  MinN(x) = x
func MinN(xs ...float64) float64 {
	min := math.Inf(1)
	for _, x := range xs {
		min = math.Min(min, x)
	}
	return min
}
