package main

import "math"

// Min returns the minimum value of an int slice
func Min(vals []int) int {
	min := math.MaxInt32
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

// Min64 returns the minimum value of an int64 slice
func Min64(vals []int64) int64 {
	var min int64
	min = math.MaxInt64
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
