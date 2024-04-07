package goutil

import "golang.org/x/exp/constraints"

// Sum returns the sum of vals.
func Sum[K constraints.Integer](vals ...K) (sum K) {
	for _, v := range vals {
		sum += v
	}
	return sum
}
