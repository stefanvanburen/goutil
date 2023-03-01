package goutil

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestDivisors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		n    int
		want []int
	}{
		{28, []int{1, 2, 4, 7, 14}},
		{12, []int{1, 2, 3, 4, 6}},
		{16, []int{1, 2, 4, 8}},
	}
	for _, test := range tests {
		if got := Divisors(test.n); !slices.Equal(got, test.want) {
			t.Errorf("Divisors(%v) = %v; want %v", test.n, got, test.want)
		}
	}
}
