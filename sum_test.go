package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-1, -2, 3}, 0},
	}

	for _, x := range tests {
		assert.Equal(t, Sum(x.in), x.out, "Should be equal")
	}
}
