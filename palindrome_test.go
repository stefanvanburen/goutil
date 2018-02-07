package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"hello", false},
		{"hehheh", true},
		{"hellex", false},
		{"helleh", true},
	}

	for _, x := range tests {
		assert.Equal(t, IsPalindrome(x.in), x.out, "Should be equal")
	}

}
