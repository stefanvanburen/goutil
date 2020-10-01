package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestMinInt(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		in  []int
		out int
	}{
		{[]int{-17, 5, 2, 18, 20, -64, 8}, -64},
		{[]int{-10, -9, -8, -6}, -10},
	}

	for _, tt := range tests {
		is.Equal(MinInt(tt.in), tt.out)
	}
}
