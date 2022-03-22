package goutil

import (
	"math"

	"golang.org/x/exp/constraints"
)

// SqrtInt determines the square root of an int.
func SqrtInt[K constraints.Integer](x K) K {
	return K(math.Sqrt(float64(x)))
}
