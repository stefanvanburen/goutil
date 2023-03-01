package goutil

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in   []int
		want int
	}{
		"positive": {[]int{1, 2, 3}, 6},
		"mixed":    {[]int{-1, -2, 3}, 0},
		"negative": {[]int{-1, -2, -3}, -6},
	}
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Sum(tc.in...); got != tc.want {
				t.Errorf("Sum(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
