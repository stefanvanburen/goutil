package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		a   string
		out bool
	}{
		{"hello", false},
		{"hehheh", true},
		{"hellex", false},
		{"helleh", true},
	}

	for _, x := range tests {
		assert.Equal(t, IsPalindrome(x.a), x.out, "Should be equal")
	}

}
