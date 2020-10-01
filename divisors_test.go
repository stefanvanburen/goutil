package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestDivisors(t *testing.T) {
	tests := []struct {
		n        int
		divisors []int
	}{
		{28, []int{1, 2, 4, 7, 14}},
		{12, []int{1, 2, 3, 4, 6}},
		{16, []int{1, 2, 4, 8}},
	}

	for _, x := range tests {
		is := is.New(t)

		is.Equal(Divisors(x.n), x.divisors)
	}
}
