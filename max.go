package goutil

import (
	"errors"
	"math"
)

// TODO: define error type enum

// MaxInt returns the maximum value of an int slice
func MaxInt(vals []int) (max int, err error) {
	if len(vals) == 0 {
		err = errors.New("Empty slice")
		return
	}
	max = math.MinInt32
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// MaxInt8 returns the maximum value of an int8 slice
func MaxInt8(vals []int8) (max int8, err error) {
	if len(vals) == 0 {
		err = errors.New("Empty slice")
		return
	}
	max = math.MinInt8
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// MaxInt16 returns the maximum value of an int16 slice
func MaxInt16(vals []int16) (max int16, err error) {
	if len(vals) == 0 {
		err = errors.New("Empty slice")
		return
	}
	max = math.MinInt16
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// MaxInt32 returns the maximum value of an int32 slice
func MaxInt32(vals []int32) (max int32, err error) {
	if len(vals) == 0 {
		err = errors.New("Empty slice")
		return
	}
	max = math.MinInt32
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// MaxInt64 returns the maximum value of an int64 slice
func MaxInt64(vals []int64) (max int64, err error) {
	if len(vals) == 0 {
		err = errors.New("Empty slice")
		return
	}
	max = math.MinInt64
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}
