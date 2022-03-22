package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestMaxInt(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{-17, 5, 2, 18, 20, -64, 8}, 20},
		{[]int{-10, -9, -8, -6}, -6},
	}

	for _, tt := range tests {
		is := is.New(t)

		is.Equal(Max(tt.in...), tt.out)
	}
}
