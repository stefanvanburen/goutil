package goutil

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"h", true}, // Single character is a palindrome
		{"hello", false},
		{"hehheh", true},
		{"hellex", false},
		{"helleh", true},
	}
	for _, x := range tests {
		if got := IsPalindrome(x.in); got != x.want {
			t.Errorf("IsPalindrome(%v) = %v, want %v", x.in, got, x.want)
		}
	}
}
