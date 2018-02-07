package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	n        int
	divisors []int
}{
	{28, []int{1, 2, 4, 7, 14}},
	{12, []int{1, 2, 3, 4, 6}},
	{16, []int{1, 2, 4, 8}},
}

func TestDivisors(t *testing.T) {
	for _, x := range tests {
		out := Divisors(x.n)
		assert.Equal(t, out, x.divisors, "Should be equal")
	}
}
