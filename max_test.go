package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{[]int{-17, 5, 2, 18, 20, -64, 8}, 20},
		{[]int{-10, -9, -8, -6}, -6},
	}

	for _, tt := range tests {
		out := MaxInt(tt.in)
		assert.Equal(t, out, tt.out)
	}
}
