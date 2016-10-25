package goutil

import "testing"

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
		c := IsPalindrome(x.a)
		if c != x.out {
			t.Errorf("IsPalindrome(%s) => %t, want %t", x.a, c, x.out)
		}
	}

}
