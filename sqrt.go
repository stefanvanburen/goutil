package main

import "math"

// Sqrt determines the square root of an int
func Sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// Sqrt64 determines the square root of an int64
func Sqrt64(x int64) int64 {
	return int64(math.Sqrt(float64(x)))
}
