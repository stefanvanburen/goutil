package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestPalindrome(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		in  string
		out bool
	}{
		{"hello", false},
		{"hehheh", true},
		{"hellex", false},
		{"helleh", true},
	}

	for _, x := range tests {
		is.Equal(IsPalindrome(x.in), x.out)
	}

}
