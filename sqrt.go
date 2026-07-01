package goutil

import "math"

// SqrtInt determines the square root of an int.
func SqrtInt[K Integer](x K) K {
	return K(math.Sqrt(float64(x)))
}
