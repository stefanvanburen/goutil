package goutil

import "golang.org/x/exp/constraints"

// Max returns the maximum value of `vals`.
func Max[K constraints.Ordered](vals ...K) K {
	var max K
	if len(vals) == 0 {
		return max
	} else if len(vals) == 1 {
		return vals[0]
	}

	max = vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max
}
