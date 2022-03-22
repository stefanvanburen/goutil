package goutil

import "golang.org/x/exp/constraints"

// Min returns the minimum value of `vals`.
func Min[K constraints.Ordered](vals ...K) K {
	var min K
	if len(vals) == 0 {
		return min
	} else if len(vals) == 1 {
		return vals[0]
	}

	min = vals[0]

	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min
}
