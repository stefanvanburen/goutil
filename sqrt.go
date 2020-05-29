package goutil

import "math"

// SqrtInt determines the square root of an int.
func SqrtInt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// SqrtInt64 determines the square root of an int64.
func SqrtInt64(x int64) int64 {
	return int64(math.Sqrt(float64(x)))
}
