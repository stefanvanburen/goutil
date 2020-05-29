package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{[]int{-17, 5, 2, 18, 20, -64, 8}, -64},
		{[]int{-10, -9, -8, -6}, -10},
	}

	for _, tt := range tests {
		out := MinInt(tt.in)
		assert.Equal(t, out, tt.out)
	}
}
