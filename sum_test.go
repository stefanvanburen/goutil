package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestSumInt(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out int
	}{
		"positive": {[]int{1, 2, 3}, 6},
		"mixed":    {[]int{-1, -2, 3}, 0},
		"negative": {[]int{-1, -2, -3}, -6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			is.Equal(Sum(tc.in...), tc.out)
		})
	}
}
