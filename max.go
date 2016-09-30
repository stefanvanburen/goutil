package main

import "math"

// Max returns the maximum value of an int slice
func Max(vals []int) int {
	max := math.MinInt32
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// Max64 returns the maximum value of an int64 slice
func Max64(vals []int64) int64 {
	var max int64
	max = math.MinInt64
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
