package goutil

import (
	"testing"
)

func TestSqrtInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   int
		want int
	}{
		{4, 2},
		{9, 3},
		{100, 10},
		{144, 12},
	}
	for _, test := range tests {
		if got := SqrtInt(test.in); got != test.want {
			t.Errorf("SqrtInt(%v) = %v, want %v", test.in, got, test.want)
		}
	}
}
