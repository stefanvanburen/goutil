package main

import "math"

// Max returns the maximum value of an int slice
func Max(vals []int) (max int) {
	max = math.MinInt32
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// Max64 returns the maximum value of an int64 slice
func Max64(vals []int64) (max int64) {
	max = int64(math.MinInt64)
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}
