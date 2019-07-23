package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInt(t *testing.T) {
	var tests = map[string]struct {
		in  []int
		out int
	}{
		"positive": {[]int{1, 2, 3}, 6},
		"mixed":    {[]int{-1, -2, 3}, 0},
		"negative": {[]int{-1, -2, -3}, -6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, SumInt(tc.in), tc.out)
		})
	}
}

func TestSumInt8(t *testing.T) {
	var tests = map[string]struct {
		in  []int8
		out int8
	}{
		"positive": {[]int8{1, 2, 3}, 6},
		"mixed":    {[]int8{-1, -2, 3}, 0},
		"negative": {[]int8{-1, -2, -3}, -6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, SumInt8(tc.in), tc.out)
		})
	}
}

func TestSumInt16(t *testing.T) {
	var tests = map[string]struct {
		in  []int16
		out int16
	}{
		"positive": {[]int16{1, 2, 3}, 6},
		"mixed":    {[]int16{-1, -2, 3}, 0},
		"negative": {[]int16{-1, -2, -3}, -6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, SumInt16(tc.in), tc.out)
		})
	}
}

func TestSumInt32(t *testing.T) {
	var tests = map[string]struct {
		in  []int32
		out int32
	}{
		"positive": {[]int32{1, 2, 3}, 6},
		"mixed":    {[]int32{-1, -2, 3}, 0},
		"negative": {[]int32{-1, -2, -3}, -6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, SumInt32(tc.in), tc.out)
		})
	}
}

func TestSumInt64(t *testing.T) {
	var tests = map[string]struct {
		in  []int64
		out int64
	}{
		"positive": {[]int64{1, 2, 3}, 6},
		"mixed":    {[]int64{-1, -2, 3}, 0},
		"negative": {[]int64{-1, -2, -3}, -6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, SumInt64(tc.in), tc.out)
		})
	}
}
