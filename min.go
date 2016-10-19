package goutil

import "math"

// Min returns the minimum value of an int slice
func Min(vals []int) (min int) {
	min = math.MaxInt32
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return
}

// Min64 returns the minimum value of an int64 slice
func Min64(vals []int64) (min int64) {
	min = math.MaxInt64
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return
}
